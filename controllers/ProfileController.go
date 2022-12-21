package controllers

import (
	"app/matchingAppProfileService/common/dataStructures"
	"app/matchingAppProfileService/common/dbInterface"
	"app/matchingAppProfileService/common/security"
	"app/matchingAppProfileService/publish"

	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

func GetAllProfiles(db *gorm.DB) gin.HandlerFunc {
	handler := func(context *gin.Context) {
		users, err := dbInterface.GetAllUsersNew(db)
		if err != nil {
			context.AbortWithStatus(http.StatusInternalServerError)
		}
		context.IndentedJSON(http.StatusOK, users)
	}

	return gin.HandlerFunc(handler)
}

func CreateProfile(db *gorm.DB, redis *redis.Client) gin.HandlerFunc {
	handler := func(context *gin.Context) {
		var newUser dataStructures.User
		if err := context.BindJSON(&newUser); err != nil {
			fmt.Println(err)
			context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "Corrupted Data send!",
			})
			return
		}
		if _, err := dbInterface.GetUserByEmail(db, newUser.Email); err == nil {
			fmt.Println("User already existing!")
			context.AbortWithStatusJSON(http.StatusConflict, gin.H{
				"error": "User with this email already exists!",
			})
			return
		}
		newUser.Confirmed = false
		newUser.Active = false
		userToReturn, errCreate := dbInterface.CreateUser(db, &newUser)
		if errCreate != nil {
			fmt.Println(errCreate)
			context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "User could not be created.",
			})
			return
		}
		signUpCode := dbInterface.CreateAndSaveSignupCode(redis, userToReturn.ID)
		publish.PublishRegister(userToReturn.ID, signUpCode)
		context.IndentedJSON(http.StatusOK, userToReturn)
	}

	return gin.HandlerFunc(handler)
}

func ActivateUser(redis *redis.Client, db *gorm.DB) gin.HandlerFunc {
	handler := func(context *gin.Context) {
		type activateObject struct {
			Code string `json:"code"`
		}
		var obj activateObject
		if err := context.BindJSON(&obj); err != nil {
			fmt.Println(err)
			context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "The activation code has to be an string",
			})
			return
		}
		id := context.Param("id")
		user, err := dbInterface.GetUserById(db, id)
		if err != nil {
			context.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		activateCode, errCode := dbInterface.GetSignUpCode(redis, user.ID)
		if errCode != nil {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "This activation code is no longer valid! Request a new one!",
			})
			return
		}
		if activateCode == "" {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "This activation code is no longer valid! Request a new one!",
			})
			return
		}
		if obj.Code != activateCode {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid activation code!",
			})
			return
		}
		ok, errActivate := dbInterface.ActivateUser(user.ID, db)
		if errActivate != nil {
			context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": errActivate,
			})
			return
		}
		if !ok {
			context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "An unkown error has occured!",
			})
			return
		}
		publish.PublishSignUp(user.ID)
		jwtString, errJWT := security.GenerateJWT(user)
		if errJWT != nil {
			context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "Could not generate access token",
			})
		}
		context.JSON(http.StatusOK, gin.H{
			"message": "User activated!",
			"token":   jwtString,
		})
	}
	return gin.HandlerFunc(handler)
}

func GetProfileById(db *gorm.DB) gin.HandlerFunc {
	handler := func(context *gin.Context) {
		id := context.Param("id")
		users, err := dbInterface.GetUserById(db, id)
		if err != nil {
			context.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		context.IndentedJSON(http.StatusOK, users)
	}
	return gin.HandlerFunc(handler)
}

func LoginUser(db *gorm.DB) gin.HandlerFunc {
	handler := func(context *gin.Context) {
		// Get the email and password from request body

		var requestBody struct {
			Email    string
			Password string
		}

		if errReqBody := context.Bind(&requestBody); errReqBody != nil {
			context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "Failed to read request body!",
			})
			return
		}

		if requestBody.Email == "" || requestBody.Password == "" {
			context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "Failed to read request body!",
			})
			return
		}

		// Get the corresponding user

		requestedUser, errReq := dbInterface.GetUserByEmail(db, requestBody.Email)

		if errReq != nil {
			context.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "Invalid email or password!",
			})
			return
		}

		// Is user activated?

		if !requestedUser.Active {
			context.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "This account has yet to be activated!",
			})
			return
		}

		// Compare sent passhash with saved passhash

		if requestedUser.Password != requestBody.Password {
			context.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "Invalid email or password!",
			})
			return
		}

		// Generate jwt token
		jwtString, err := security.GenerateJWT(requestedUser)
		if err != nil {
			context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
		}
		// Send the jwt token
		context.IndentedJSON(http.StatusOK, gin.H{
			"token": jwtString,
		})
	}
	return gin.HandlerFunc(handler)
}

func DeleteUser(db *gorm.DB) gin.HandlerFunc {
	handler := func(context *gin.Context) {
		var toFind struct {
			Email string
		}
		errExtract := context.Bind(&toFind)
		if errExtract != nil {
			context.AbortWithStatus(http.StatusBadRequest)
			return
		}

		userToDelete, errFind := dbInterface.GetUserByEmail(db, toFind.Email)
		if errFind != nil {
			context.AbortWithError(http.StatusNotFound, errFind)
			return
		}

		if errDelete := dbInterface.DeleteUser(db, userToDelete); errDelete != nil {
			context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": errDelete,
			})
			return
		}

		context.JSON(http.StatusOK, gin.H{
			"message": "User deleted!",
		})
	}
	return gin.HandlerFunc(handler)
}

func UpdateUser(db *gorm.DB) gin.HandlerFunc {
	handler := func(context *gin.Context) {
		var newData *dataStructures.User
		var userId = context.Param("id")
		errBind := context.BindJSON(&newData)
		if errBind != nil {
			context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": errBind,
			})
			return
		}

		updatedUser, errUpdate := dbInterface.UpdateUser(db, userId, newData)
		if errUpdate != nil {
			context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": errBind,
			})
			return
		}
		context.JSON(http.StatusOK, updatedUser)

	}
	return gin.HandlerFunc(handler)
}

func ValidateJWT(db *gorm.DB) gin.HandlerFunc {
	handler := func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "Authorization valid",
		})
	}
	return gin.HandlerFunc(handler)
}

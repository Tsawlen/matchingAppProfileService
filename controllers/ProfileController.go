package controllers

import (
	"app/matchingAppProfileService/common/dataStructures"
	"app/matchingAppProfileService/common/dbInterface"
	"app/matchingAppProfileService/common/security"

	"fmt"
	"net/http"

	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
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

func CreateProfile(db *gorm.DB) gin.HandlerFunc {
	handler := func(context *gin.Context) {
		var newUser dataStructures.User
		if err := context.BindJSON(&newUser); err != nil {
			fmt.Println(err)
			context.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		if _, err := dbInterface.GetUserByEmail(db, newUser.Email); err == nil {
			fmt.Println("User already existing!")
			context.AbortWithStatusJSON(http.StatusConflict, gin.H{
				"error": "User with this email already exists!",
			})
			return
		}
		userToReturn, errCreate := dbInterface.CreateUser(db, &newUser)
		if errCreate != nil {
			fmt.Println(errCreate)
			context.AbortWithError(http.StatusInternalServerError, errCreate)
			return
		}
		context.IndentedJSON(http.StatusOK, userToReturn)
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

		// Compare sent passhash with saved passhash

		if requestedUser.Password != requestBody.Password {
			context.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "Invalid email or password!",
			})
			return
		}

		// Generate jwt token

		token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
			"sub": requestedUser.ID,
			"exp": time.Now().Add(3 * time.Hour).Unix(),
		})

		// Get RSA private key

		key, errGetKey := security.GetPrivateToken()

		if errGetKey != nil {
			fmt.Println(errGetKey)
			context.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "Failed to create token",
			})
			return
		}

		// Sign the token
		tokenString, errSign := token.SignedString(key)

		if errSign != nil {
			fmt.Println(errSign)
			context.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "Failed to create token",
			})
			return
		}
		// Send the jwt token
		context.IndentedJSON(http.StatusOK, gin.H{
			"token": tokenString,
		})
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
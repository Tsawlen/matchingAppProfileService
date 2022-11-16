package query

import (
	"app/matchingAppProfileService/common/dataStructures"
	"app/matchingAppProfileService/common/dbInterface"
	"app/matchingAppProfileService/common/mockData"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func queryProfiles(id string) (*dataStructures.Profile, error) {
	for counter, value := range mockData.ProfileData {
		if value.Id == id {
			return &mockData.ProfileData[counter], nil
		}
	}
	return &dataStructures.Profile{}, errors.New("profile not found!")
}

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
		}
		userToReturn, errCreate := dbInterface.CreateUser(db, &newUser)
		if errCreate != nil {
			fmt.Println(errCreate)
			context.AbortWithError(http.StatusInternalServerError, errCreate)
		}
		context.IndentedJSON(http.StatusOK, userToReturn)
	}

	return gin.HandlerFunc(handler)
}

func GetAllProfilesMock(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, mockData.ProfileData)
}

func GetProfileById(db *gorm.DB) gin.HandlerFunc {
	handler := func(context *gin.Context) {
		id := context.Param("id")
		users, err := dbInterface.GetUserById(db, id)
		if err != nil {
			context.AbortWithStatus(http.StatusInternalServerError)
		}
		context.IndentedJSON(http.StatusOK, users)
	}

	return gin.HandlerFunc(handler)
}

/*func GetProfileById(context *gin.Context) {
	id := context.Param("id")
	searchedProfile, error := queryProfiles(id)
	if error != nil {
		context.AbortWithStatus(http.StatusNotFound)
	}
	context.IndentedJSON(http.StatusOK, searchedProfile)
}*/

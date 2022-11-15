package query

import (
	"app/matchingAppProfileService/common/crud"
	"app/matchingAppProfileService/common/dataStructures"
	"app/matchingAppProfileService/common/mockData"
	"database/sql"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func queryProfiles(id string) (*dataStructures.Profile, error) {
	for counter, value := range mockData.ProfileData {
		if value.Id == id {
			return &mockData.ProfileData[counter], nil
		}
	}
	return &dataStructures.Profile{}, errors.New("profile not found!")
}

func GetAllProfilesDB(db *sql.DB) gin.HandlerFunc {
	handler := func(context *gin.Context) {
		users, err := crud.GetAllUsers(db)
		if err != nil {
			context.AbortWithStatus(http.StatusInternalServerError)
		}
		context.IndentedJSON(http.StatusOK, users)
	}

	return gin.HandlerFunc(handler)
}

func GetAllProfiles(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, mockData.ProfileData)
}

func GetProfileById(context *gin.Context) {
	id := context.Param("id")
	searchedProfile, error := queryProfiles(id)
	if error != nil {
		context.AbortWithStatus(http.StatusNotFound)
	}
	context.IndentedJSON(http.StatusOK, searchedProfile)
}

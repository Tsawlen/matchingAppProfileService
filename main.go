package main

import (
	"net/http"

	"app/matchingAppProfileService/common/dataStructures"
	"app/matchingAppProfileService/common/mockData"
	"app/matchingAppProfileService/query"

	"github.com/gin-gonic/gin"
)

func addProfile(context *gin.Context) {
	var newProfile dataStructures.Profile

	if err := context.BindJSON(&newProfile); err != nil {
		return
	}

	mockData.ProfileData = append(mockData.ProfileData, newProfile)

	context.IndentedJSON(http.StatusCreated, newProfile)
}

func main() {
	router := gin.Default()
	router.GET("/profile", query.GetAllProfiles)
	router.GET("/profile/:id", query.GetProfileById)
	router.PUT("/profile", addProfile)
	router.Run("localhost:8080")
}

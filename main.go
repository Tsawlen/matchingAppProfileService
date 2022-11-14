package main

import (
	"errors"
	"fmt"
	"net/http"

	"app/matchingAppProfileService/common/dataStructures"
	"app/matchingAppProfileService/common/database"
	"app/matchingAppProfileService/common/mockData"

	"app/matchingAppProfileService/query"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
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
	db := database.InitalizeConnection()
	router := gin.Default()
	router.GET("/profile", query.GetAllProfiles)
	router.GET("/profile/:id", query.GetProfileById)
	router.PUT("/profile", addProfile)
	router.Run("0.0.0.0:8080")
	_, err := db.Query("SELECT * FROM user")

	if err != nil {
		fmt.Println(err)
		panic(errors.New("Error querying data"))
	}
	/*
		for data.Next() {
			var userData user

			err = data.Scan(&userData.id, &userData.city, &userData.email, &userData.first_name, &userData.name, &userData.number, &userData.password, &userData.street, &userData.username)

			if err != nil {
				fmt.Println(err)
				panic(errors.New("Error finding data"))
			}

			fmt.Println(userData.username)
		}
	*/
}

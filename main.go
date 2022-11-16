package main

import (
	"database/sql"
	"net/http"

	"app/matchingAppProfileService/common/dataStructures"
	"app/matchingAppProfileService/common/database"
	"app/matchingAppProfileService/common/initializer"
	"app/matchingAppProfileService/common/mockData"

	"app/matchingAppProfileService/query"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
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
	dbChannel := make(chan *sql.DB)
	gdbChannel := make(chan *gorm.DB)
	go database.InitalizeConnection(dbChannel, gdbChannel)
	go initializer.LoadEnvVariables()

	db := <-dbChannel
	gdb := <-gdbChannel

	defer db.Close()

	router := gin.Default()
	router.GET("/profile", query.GetAllProfiles(gdb))
	router.GET("/profile/:id", query.GetProfileById(gdb))
	router.GET("/skill", query.GetAllSkills(gdb))
	router.PUT("/profile", query.CreateProfile(gdb))
	router.PUT("/skill", query.CreateSkill(gdb))
	router.Run("0.0.0.0:8080")

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

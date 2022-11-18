package main

import (
	"database/sql"

	"app/matchingAppProfileService/common/database"
	"app/matchingAppProfileService/common/initializer"
	"app/matchingAppProfileService/middleware"

	"app/matchingAppProfileService/controllers"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dbChannel := make(chan *sql.DB)
	gdbChannel := make(chan *gorm.DB)
	go database.InitalizeConnection(dbChannel, gdbChannel)
	go initializer.LoadEnvVariables()

	db := <-dbChannel
	gdb := <-gdbChannel

	defer db.Close()

	router := gin.Default()
	// Get Requests
	router.GET("/profile", middleware.RequireAuth, controllers.GetAllProfiles(gdb))
	router.GET("/profile/:id", middleware.RequireAuth, controllers.GetProfileById(gdb))
	router.GET("/skill", controllers.GetAllSkills(gdb))
	router.GET("/validate", middleware.RequireAuth, controllers.ValidateJWT(gdb))
	router.GET("/skill/:id/users", middleware.RequireAuth, controllers.GetUsersBySkill(gdb))

	// Put Requests
	router.PUT("/signUp", controllers.CreateProfile(gdb))
	router.PUT("/skill", middleware.RequireAuth, controllers.CreateSkill(gdb))
	router.PUT("/login", controllers.LoginUser(gdb))

	// Update Requests
	router.PUT("/profile/:id", middleware.RequireAuth, controllers.UpdateUser(gdb))

	// Delete Requests
	router.DELETE("/profile", middleware.RequireAuth, controllers.DeleteUser(gdb))
	router.DELETE("/skill/:id", middleware.RequireAuth, controllers.DeleteSkill(gdb))

	router.Run("0.0.0.0:8080")
}

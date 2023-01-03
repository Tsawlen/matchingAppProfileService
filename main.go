package main

import (
	//"flag"

	"database/sql"

	"app/matchingAppProfileService/common/database"
	"app/matchingAppProfileService/common/initializer"
	"app/matchingAppProfileService/middleware"

	"app/matchingAppProfileService/controllers"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

func main() {

	/*certificateFile := flag.String("certfile", "certificates/cert.pem", "certificate PEM file")
	keyFile := flag.String("keyfile", "certificates/key.pem", "key PEM file")
	flag.Parse()*/

	dbChannel := make(chan *sql.DB)
	gdbChannel := make(chan *gorm.DB)
	redisChannel := make(chan *redis.Client)
	finishedInit := make(chan bool)

	go initializer.LoadEnvVariables(finishedInit)
	<-finishedInit
	go database.InitalizeConnection(dbChannel, gdbChannel)
	go database.InitRedis(redisChannel)

	db := <-dbChannel
	gdb := <-gdbChannel
	redis := <-redisChannel

	defer db.Close()

	router := gin.Default()
	// Get Requests
	router.GET("/profile", middleware.RequireAuth, controllers.GetAllProfiles(gdb))
	router.GET("/profile/:id", middleware.RequireAuth, controllers.GetProfileById(gdb))
	router.GET("/skill", controllers.GetAllSkills(gdb))
	router.GET("/validate", middleware.RequireAuth, controllers.ValidateJWT(gdb))
	router.GET("/skill/:id/users", middleware.RequireAuth, controllers.GetUsersBySkill(gdb))
	router.GET("/skillLevel/:id", middleware.RequireAuth, controllers.GetSkillLevelById(gdb))
	router.GET("/skillLevel", middleware.RequireAuth, controllers.GetAllSkillLevels(gdb))

	// Put Requests
	router.PUT("/signUp", controllers.CreateProfile(gdb, redis))
	router.PUT("/skill", middleware.RequireAuth, controllers.CreateSkill(gdb))
	router.PUT("/login", controllers.LoginUser(gdb))
	router.PUT("/activate/:id", controllers.ActivateUser(redis, gdb))
	router.PUT("/skillLevel", middleware.RequireAuth, controllers.CreateSkillLevel(gdb))

	// Update Requests
	router.PUT("/profile/:id", middleware.RequireAuth, controllers.UpdateUser(gdb))

	// Delete Requests
	router.DELETE("/profile", middleware.RequireAuth, controllers.DeleteUser(gdb))
	router.DELETE("/skill/:id", middleware.RequireAuth, controllers.DeleteSkill(gdb))

	router.Run("0.0.0.0:8080")
}

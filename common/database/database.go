package database

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

	"app/matchingAppProfileService/common/dataStructures"

	"github.com/go-redis/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitalizeConnection(dbChannel chan *sql.DB, gdbChannel chan *gorm.DB) *sql.DB {
	dsn := "root:root@tcp(" + os.Getenv("MYSQL_HOST") + ")/golang_docker?parseTime=true"
	gDb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
		panic(errors.New("Error connecting to mysql"))
	}
	fmt.Println("Database connected!")

	db, errGetDb := gDb.DB()

	if errGetDb != nil {
		fmt.Println(err)
		panic(errors.New("Error getting DB from gorm"))
	}

	errPing := db.Ping()
	if errPing != nil {
		fmt.Println(errPing)
	}
	setupDatabase(gDb)
	//addMockData(gDb)
	dbChannel <- db
	gdbChannel <- gDb
	return db
}

func setupDatabase(db *gorm.DB) {
	db.AutoMigrate(&dataStructures.User{})
	db.AutoMigrate(&dataStructures.Skill{})
	db.AutoMigrate(&dataStructures.City{})
}

// Redis

func InitRedis(redisChannel chan *redis.Client) {
	dbInt, errInt := strconv.Atoi(os.Getenv("REDIS_DB"))
	if errInt != nil {
		log.Fatal("REDIS_DB needs to be an integer value")
	}
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_IP"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       dbInt,
	})
	if err := client.Ping().Err(); err != nil {
		log.Fatal("No connection to Redis possible")
	}
	redisChannel <- client

}

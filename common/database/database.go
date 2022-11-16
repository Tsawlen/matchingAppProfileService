package database

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"app/matchingAppProfileService/common/dataStructures"
	"app/matchingAppProfileService/common/dbInterface"
	"app/matchingAppProfileService/common/mockData"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitalizeConnection(dbChannel chan *sql.DB, gdbChannel chan *gorm.DB) *sql.DB {
	dsn := "root:root@tcp(database:3306)/golang_docker?parseTime=true"
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
	addMockData(gDb)
	dbChannel <- db
	gdbChannel <- gDb
	return db
}

func setupDatabase(db *gorm.DB) {
	db.AutoMigrate(&dataStructures.User{})
	db.AutoMigrate(&dataStructures.Skill{})
}

func addMockData(db *gorm.DB) {
	_, err := dbInterface.CreateUser(db, &mockData.UserData[0])
	if err != nil {
		log.Fatal(err)
	}
}

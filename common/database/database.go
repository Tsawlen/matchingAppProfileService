package database

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"app/matchingAppProfileService/common/crud"
	"app/matchingAppProfileService/common/mockData"

	_ "github.com/go-sql-driver/mysql"
)

func InitalizeConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(database:3306)/golang_docker")
	if err != nil {
		fmt.Println(err)
		panic(errors.New("Error connecting to mysql"))
	}
	defer db.Close()

	fmt.Println("Database connected!")

	errPing := db.Ping()
	if errPing != nil {
		fmt.Println(errPing)
	}
	createUserTable(db)
	addMockData(db)
	return db
}

func createUserTable(db *sql.DB) error {
	fmt.Println("Creating table...")
	query := "CREATE TABLE IF NOT EXISTS users(id int primary key AUTO_INCREMENT, city varchar(255), email varchar(255), first_name varchar(255), name varchar(255), password varchar(255), street varchar(255), houseNumber varchar(255), username varchar(255), created_at datetime default CURRENT_TIMESTAMP, updated_at datetime default CURRENT_TIMESTAMP)"
	fmt.Println("Sending Command!")
	res, err := db.Exec(query)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("Command sended!")
	rows, err := res.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Printf("Rows affected when creating table: %d\n", rows)
	return nil
}

func addMockData(db *sql.DB) {
	err := crud.AddUser(&mockData.UserData[0], db)
	if err != nil {
		log.Fatal(err)
	}
}

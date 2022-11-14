package database

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func InitalizeConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/golang_docker")
	if err != nil {
		fmt.Println(err)
		panic(errors.New("Error connecting to mysql"))
	}
	defer db.Close()

	return db
}

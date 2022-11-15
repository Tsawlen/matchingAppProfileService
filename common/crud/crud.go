package crud

import (
	"database/sql"
	"log"

	"app/matchingAppProfileService/common/dataStructures"

	_ "github.com/go-sql-driver/mysql"
)

func AddUser(user *dataStructures.User, db *sql.DB) error {

	statement, err := db.Prepare("INSERT INTO `users`(`id`,`city`,`email`,`first_name`,`name`,`password`,`street`,`houseNumber`,`username`)VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
		return err
	}

	_, errInsert := statement.Exec(&user.Id, &user.City, &user.Email, &user.First_name, &user.Name, &user.Password, &user.Street, &user.HouseNumber, &user.Username)

	if errInsert != nil {
		log.Fatal(errInsert)
		return errInsert
	}

	return nil

}

package crud

import (
	"database/sql"
	"fmt"
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

func GetAllUsers(db *sql.DB) (*[]dataStructures.User, error) {
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var users []dataStructures.User
	for rows.Next() {
		var user dataStructures.User
		if errLine := rows.Scan(&user.Id, &user.City, &user.Email, &user.First_name, &user.Name, &user.Password, &user.Street, &user.HouseNumber, &user.Username, &user.Created_at, &user.Updated_at); errLine != nil {
			fmt.Println(errLine)
			return nil, errLine
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return &users, nil
}

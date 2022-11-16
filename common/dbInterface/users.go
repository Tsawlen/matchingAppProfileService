package dbInterface

import (
	"fmt"

	"app/matchingAppProfileService/common/dataStructures"

	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB, user *dataStructures.User) (*dataStructures.User, error) {
	result := db.Create(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func GetAllUsersNew(db *gorm.DB) (*[]dataStructures.User, error) {
	var users []dataStructures.User

	err := db.Model(&dataStructures.User{}).Preload("Skills").Find(&users).Error

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &users, nil
}

func GetUserById(db *gorm.DB, id string) (*[]dataStructures.User, error) {
	var users []dataStructures.User

	err := db.Model(&dataStructures.User{}).Preload("Skills").Where("id=?", id).Find(&users).Error

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &users, nil
}

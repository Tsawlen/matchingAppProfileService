package dbInterface

import (
	"errors"
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

func GetUserById(db *gorm.DB, id string) (*dataStructures.User, error) {
	var users dataStructures.User

	err := db.Model(&dataStructures.User{}).Preload("Skills").Where("id=?", id).Find(&users).Error

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &users, nil
}

func GetUserByEmail(db *gorm.DB, email string) (*dataStructures.User, error) {
	var user dataStructures.User

	err := db.Model(&dataStructures.User{}).Where("Email=?", email).First(&user).Error

	if err != nil {
		fmt.Println(err)
		return nil, errors.New("No User found for this email")
	}

	return &user, nil
}

func UpdateUser(db *gorm.DB, userId string, newData *dataStructures.User) (*dataStructures.User, error) {
	userToUpdate, errFind := GetUserById(db, userId)
	if errFind != nil {
		return nil, errFind
	}

	changedUser := updateValuesForUser(userToUpdate, newData, db)

	result := db.Save(&changedUser)

	if result.Error != nil {
		return nil, result.Error
	}

	return changedUser, nil
}

func DeleteUser(db *gorm.DB, user *dataStructures.User) error {
	errAssocClear := db.Model(&user).Association("Skills").Clear()

	if errAssocClear != nil {
		return errAssocClear
	}

	result := db.Delete(&user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// Helper Functions

func updateValuesForUser(oldUser *dataStructures.User, newUser *dataStructures.User, db *gorm.DB) *dataStructures.User {
	if newUser.Skills != nil {
		errAssocClear := db.Model(&oldUser).Association("Skills").Clear()
		if errAssocClear != nil {
			fmt.Println("Could not delete Skills!")
		}
	}
	oldUser.City = newUser.City
	oldUser.First_name = newUser.First_name
	oldUser.Name = newUser.Name
	oldUser.Street = newUser.Street
	oldUser.HouseNumber = newUser.HouseNumber
	oldUser.Gender = newUser.Gender
	oldUser.Skills = newUser.Skills
	return oldUser
}

package dbInterface

import (
	"errors"
	"fmt"

	"app/matchingAppProfileService/common/dataStructures"

	"gorm.io/gorm"
)

func CreateSkill(db *gorm.DB, skill *dataStructures.Skill) (*dataStructures.Skill, error) {
	result := db.Create(&skill)
	if result.Error != nil {
		fmt.Println(result.Error)
		return nil, result.Error
	}
	return skill, nil
}

func GetAllSkills(db *gorm.DB) (*[]dataStructures.Skill, error) {
	var skills []dataStructures.Skill
	result := db.Find(&skills)

	if result.Error != nil {
		fmt.Println(result.Error)
		return nil, result.Error
	}
	return &skills, nil
}

func GetSkillById(db *gorm.DB, id string) (*dataStructures.Skill, error) {
	var skill dataStructures.Skill
	result := db.Where("id = ?", id).First(&skill)

	if result.Error != nil {
		fmt.Println(result.Error)
		return nil, result.Error
	}
	return &skill, nil
}

func GetUsersBySkill(db *gorm.DB, id string) ([]*dataStructures.User, error) {
	var skill *dataStructures.Skill

	err := db.Model(&dataStructures.Skill{}).Preload("Users").Where("id=?", id).First(&skill).Error
	if err != nil {
		return nil, err
	}
	if len(skill.Users) <= 0 {
		return nil, errors.New("No Users for this skill found!")
	}
	return skill.Users, nil
}

func DeleteSkill(db *gorm.DB, skill *dataStructures.Skill) error {
	errAssocClear := db.Model(&skill).Association("Users").Clear()
	if errAssocClear != nil {
		return errAssocClear
	}
	result := db.Delete(&skill)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

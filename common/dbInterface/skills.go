package dbInterface

import (
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

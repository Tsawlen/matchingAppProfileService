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
	result := db.Model(&dataStructures.Skill{}).Preload("SkillLevel").Find(&skills) /*db.Find(&skills).Preload(clause.Associations)*/

	if result.Error != nil {
		fmt.Println(result.Error)
		return nil, result.Error
	}
	return &skills, nil
}

func GetSkillById(db *gorm.DB, id string) (*dataStructures.Skill, error) {
	var skill dataStructures.Skill
	result := db.Model(&dataStructures.Skill{}).Preload("SkillLevel").Where("id=?", id).Find(&skill)

	if result.Error != nil {
		fmt.Println(result.Error)
		return nil, result.Error
	}
	return &skill, nil
}

func GetUsersBySkill(db *gorm.DB, id string) ([]*dataStructures.User, error) {
	skillCache, errSkill := GetSkillById(db, id)
	if errSkill != nil {
		return nil, errSkill
	}
	allSkillLevels, errSkillLevels := GetAllSkillLevels(db)
	if errSkillLevels != nil {
		return nil, errSkillLevels
	}
	arrSkillId := []int{}
	for _, data := range *allSkillLevels {
		if data.ID >= skillCache.SkillLevel.ID {
			arrSkillId = append(arrSkillId, data.ID)
		}
	}
	var skills *[]dataStructures.Skill

	err := db.Model(&dataStructures.Skill{}).Preload("UsersAchieved").Where("name=? AND (skill_identifier) IN ?", skillCache.Name, arrSkillId).Find(&skills).Error
	if err != nil {
		return nil, err
	}
	users := []*dataStructures.User{}
	for _, skill := range *skills {
		for _, user := range skill.UsersAchieved {
			if !checkForExisting(users, user) {
				users = append(users, user)
			}
		}
	}
	if len(users) <= 0 {
		return nil, errors.New("No Users for this skill found!")
	}
	return users, nil
}

func DeleteSkill(db *gorm.DB, skill *dataStructures.Skill) error {
	errAssocClearSearch := db.Model(&skill).Association("UsersSearching").Clear()
	errAssocClearAchieved := db.Model(&skill).Association("UsersAchieved").Clear()
	if errAssocClearSearch != nil {
		return errAssocClearSearch
	}
	if errAssocClearAchieved != nil {
		return errAssocClearAchieved
	}
	result := db.Delete(&skill)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Helper

func checkForExisting(users []*dataStructures.User, user *dataStructures.User) bool {
	for _, userIn := range users {
		if userIn == user {
			return true
		}
	}
	return false
}

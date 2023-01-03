package dbInterface

import (
	"app/matchingAppProfileService/common/dataStructures"

	"gorm.io/gorm"
)

func CreateSkillLevel(db *gorm.DB, skillLevel *dataStructures.SkillLevel) (bool, error) {
	result := db.Create(&skillLevel)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

func GetAllSkillLevels(db *gorm.DB) (*[]dataStructures.SkillLevel, error) {
	var skillLevels *[]dataStructures.SkillLevel
	result := db.Find(&skillLevels)
	if result.Error != nil {
		return nil, result.Error
	}
	return skillLevels, nil
}

func GetSkillLevelById(db *gorm.DB, id int) (*dataStructures.SkillLevel, error) {
	var skillLevel *dataStructures.SkillLevel
	result := db.Where("id=?", id).First(&skillLevel)
	if result.Error != nil {
		return &dataStructures.SkillLevel{}, result.Error
	}
	return skillLevel, nil
}

func GetSkillLevelByName(db *gorm.DB, name string) (*dataStructures.SkillLevel, error) {
	var skillLevel *dataStructures.SkillLevel
	result := db.Where("name=?", name).First(&skillLevel)
	if result.Error != nil {
		return &dataStructures.SkillLevel{}, result.Error
	}
	return skillLevel, nil
}

func DeleteSkillLevel(db *gorm.DB, id int) (bool, error) {
	result := db.Where("id=?", id).Delete(&dataStructures.SkillLevel{})
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

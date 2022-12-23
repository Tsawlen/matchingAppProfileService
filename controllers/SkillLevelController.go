package controllers

import (
	"app/matchingAppProfileService/common/dataStructures"
	"app/matchingAppProfileService/common/dbInterface"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateSkillLevel(db *gorm.DB) gin.HandlerFunc {
	handler := func(context *gin.Context) {
		var newSkillLevel *dataStructures.SkillLevel
		if err := context.Bind(&newSkillLevel); err != nil {
			context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "Please provide a valid Skill Level object!",
			})
			return
		}
		ok, err := dbInterface.CreateSkillLevel(db, newSkillLevel)
		if !ok {
			context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
			return
		}
		context.IndentedJSON(http.StatusCreated, gin.H{
			"message": "Skill level created!",
		})
	}
	return gin.HandlerFunc(handler)
}

func GetAllSkillLevels(db *gorm.DB) gin.HandlerFunc {
	handler := func(context *gin.Context) {
		skillLevels, err := dbInterface.GetAllSkillLevels(db)
		if err != nil {
			context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
		}
		context.IndentedJSON(http.StatusOK, skillLevels)
	}
	return gin.HandlerFunc(handler)
}

func GetSkillLevelById(db *gorm.DB) gin.HandlerFunc {
	handler := func(context *gin.Context) {
		skillLevelId := context.Param("id")
		skillLevelUUID, errParse := strconv.Atoi(skillLevelId)
		if errParse != nil {
			context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "Please provide a valid id!",
			})
			return
		}
		skillLevel, err := dbInterface.GetSkillLevelById(db, skillLevelUUID)
		if err != nil {
			context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
			return
		}
		var badSkillLevel dataStructures.SkillLevel
		if skillLevel == &badSkillLevel {
			context.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"message": "Could not find Skill Level with id: " + skillLevelId,
			})
			return
		}
		context.IndentedJSON(http.StatusOK, skillLevel)
	}
	return gin.HandlerFunc(handler)
}

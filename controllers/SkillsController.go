package controllers

import (
	"app/matchingAppProfileService/common/dataStructures"
	"app/matchingAppProfileService/common/dbInterface"
	"fmt"

	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAllSkills(db *gorm.DB) gin.HandlerFunc {
	handler := func(context *gin.Context) {
		skills, err := dbInterface.GetAllSkills(db)
		if err != nil {
			context.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		context.IndentedJSON(http.StatusOK, &skills)
	}
	return gin.HandlerFunc(handler)
}

func CreateSkill(db *gorm.DB) gin.HandlerFunc {
	handler := func(context *gin.Context) {
		var newSkill dataStructures.Skill
		if err := context.BindJSON(&newSkill); err != nil {
			fmt.Println(err)
			context.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		skill, errCreate := dbInterface.CreateSkill(db, &newSkill)
		if errCreate != nil {
			fmt.Println(errCreate)
			context.AbortWithError(http.StatusInternalServerError, errCreate)
			return
		}
		context.IndentedJSON(http.StatusCreated, &skill)
	}
	return gin.HandlerFunc(handler)
}

func DeleteSkill(db *gorm.DB) gin.HandlerFunc {
	handler := func(context *gin.Context) {
		skillId := context.Param("id")

		skillToDelete, findErr := dbInterface.GetSkillById(db, skillId)
		if findErr != nil {
			context.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error": "Skill with id " + skillId + " not found!",
			})
			return
		}

		deleteErr := dbInterface.DeleteSkill(db, skillToDelete)
		if deleteErr != nil {
			context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": deleteErr,
			})
			return
		}

		context.JSON(http.StatusOK, gin.H{
			"message": "Skill with id " + skillId + " deleted!",
		})

	}
	return gin.HandlerFunc(handler)
}

func GetUsersBySkill(db *gorm.DB) gin.HandlerFunc {
	handler := func(context *gin.Context) {
		id := context.Param("id")
		users, err := dbInterface.GetUsersBySkill(db, id)
		if err != nil {
			context.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error": err,
			})
			return
		}
		context.JSON(http.StatusFound, users)
	}
	return gin.HandlerFunc(handler)
}

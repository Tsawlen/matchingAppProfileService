package query

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
		}
		skill, errCreate := dbInterface.CreateSkill(db, &newSkill)
		if errCreate != nil {
			fmt.Println(errCreate)
			context.AbortWithError(http.StatusInternalServerError, errCreate)
		}
		context.IndentedJSON(http.StatusCreated, &skill)
	}
	return gin.HandlerFunc(handler)
}

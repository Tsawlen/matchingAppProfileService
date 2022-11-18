package dataStructures

import (
	"time"
)

type User struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdatedTime"`
	City        string    `json:"city"`
	Email       string    `json:"email"`
	First_name  string    `json:"first_name"`
	Name        string    `json:"name"`
	Password    string    `json:"password"`
	Street      string    `json:"street"`
	HouseNumber string    `json:"houseNumber"`
	Username    string    `json:"username"`
	Gender      string    `json:"gender"`
	Skills      []*Skill  `json:"skills" gorm:"many2many:user_skills;contraint:OnDelete:NONE"`
}

type Skill struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdatedTime"`
	Name      string    `json:"name"`
	Level     string    `json:"level"`
	Users     []*User   `json:"users" gorm:"many2many:user_skills;contraint:OnDelete:CASCADE"`
}

type RemoveSkill struct {
	UserId   string   `json:"userid"`
	SkillIds []string `json:"skill_ids"`
}

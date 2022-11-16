package dataStructures

import (
	"time"
)

type Profile struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type User struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdatedTime"`
	DeletedAt   time.Time `json:"deleted_at" gorm:"default:null"`
	City        string    `json:"city"`
	Email       string    `json:"email"`
	First_name  string    `json:"first_name"`
	Name        string    `json:"name"`
	Password    string    `json:"password"`
	Street      string    `json:"street"`
	HouseNumber string    `json:"houseNumber"`
	Username    string    `json:"username"`
	Skills      []*Skill  `json:"skills" gorm:"many2many:user_skills"`
}

type Skill struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdatedTime"`
	DeletedAt time.Time `json:"deleted_at" gorm:"default:null"`
	Name      string    `json:"name"`
	Level     string    `json:"level"`
	Users     []*User   `json:"users" gorm:"many2many:user_skills"`
}

package dataStructures

import (
	"time"
)

type User struct {
	ID              uint      `json:"id" gorm:"primaryKey"`
	CreatedAt       time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt       time.Time `json:"updated_at" gorm:"autoUpdatedTime"`
	First_name      string    `json:"firstName"`
	Name            string    `json:"name"`
	Gender          uint      `json:"gender"`
	Username        string    `json:"username"`
	Email           string    `json:"email"`
	Street          string    `json:"street"`
	HouseNumber     string    `json:"houseNumber"`
	TelephoneNumber string    `json:"telephoneNumber"`
	Price           float64   `json:"price"`
	ProfilPicture   []byte    `json:"profilePicture"`
	Confirmed       bool      `json:"confirmed"`
	Active          bool      `json:"active"`
	Password        string    `json:"password"`
	SearchedSkills  []*Skill  `json:"searchedSkills" gorm:"many2many:user_searchedSkills"`
	AchievedSkills  []*Skill  `json:"achievedSkills" gorm:"many2many:user_achievedSkills"`
	CityIdentifier  int
	City            *City `json:"city" gorm:"foreignKey:CityIdentifier"`
}

type Skill struct {
	ID             uint      `json:"id" gorm:"primaryKey"`
	CreatedAt      time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt      time.Time `json:"updated_at" gorm:"autoUpdatedTime"`
	Name           string    `json:"name"`
	Level          string    `json:"level"`
	UsersSearching []*User   `json:"usersSearching" gorm:"many2many:user_searchedSkills"`
	UsersAchieved  []*User   `json:"usersAchieved" gorm:"many2many:user_achievedSkills"`
}

type RemoveSkill struct {
	UserId   string   `json:"userid"`
	SkillIds []string `json:"skill_ids"`
}

type City struct {
	PLZ       uint      `json:"plz" gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdatedTime"`
	Place     string    `json:"place"`
}

// DAO

type UserFilter struct {
	Gender uint    `json:"gender"`
	Price  float64 `json:"price"`
}

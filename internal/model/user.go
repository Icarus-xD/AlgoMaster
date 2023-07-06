package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Login      string `gorm:"unique;not null" json:"login"`
	FirstName  string    `gorm:"not null" json:"first_name"`
	MiddleName string    `gorm:"not null" json:"middle_name"`
	LastName   string    `gorm:"not null" json:"last_name"`
	Group      string    `gorm:"not null" json:"group"`
	Email      string    `gorm:"not null" json:"email"`
	Debt       int       `gorm:"not null" json:"debt"`
}

func (User) TableName() string {
	return "users"
}
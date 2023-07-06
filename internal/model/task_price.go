package model

import "gorm.io/gorm"

type TaskPrice struct {
	gorm.Model
	Type string `gorm:"unique;not null"`
	Price int `gorm:"not null"`
}

func (TaskPrice) TableName() string {
	return "task_prices"
}
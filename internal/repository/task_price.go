package repository

import (
	"github.com/Icarus-xD/AlgoMaster/internal/model"
	"gorm.io/gorm"
)

type TaskPriceRepo struct {
	db *gorm.DB
}

func NewTaskPriceRepo(db *gorm.DB) *TaskPriceRepo {
	return &TaskPriceRepo{
		db: db,
	}
}

func (r *TaskPriceRepo) GetPrice(t string) (*model.TaskPrice, error) {
	var price model.TaskPrice

	err := r.db.Where("type = ?", t).First(&price).Error
	if err != nil {
		return nil, err
	}

	return &price, nil
}
package repository

import (
	"errors"

	"github.com/Icarus-xD/AlgoMaster/internal/dto"
	"github.com/Icarus-xD/AlgoMaster/internal/model"
	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (r *UserRepo) Create(login string, payload dto.CreateUserDTO) (*model.User, error) {
	user := model.User{
		Login: login,
		FirstName: payload.FirstName,
		MiddleName: payload.MiddleName,
		LastName: payload.LastName,
		Group: payload.Group,
		Email: payload.Email,
		Debt: 0,
	}

	err := r.db.Create(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepo) UpdateDebt(login string, price, maxDebt int) (*model.User, error) {
	var user model.User

	err := r.db.Where("login = ?", login).First(&user).Error
	if err != nil {
		return nil, err
	}

	if user.Debt >= maxDebt {
		return nil, errors.New("accumulated maximum debt")
	}

	user.Debt += price
	err = r.db.Save(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
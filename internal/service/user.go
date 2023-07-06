package service

import (
	"github.com/Icarus-xD/AlgoMaster/internal/dto"
	"github.com/Icarus-xD/AlgoMaster/internal/model"
	"github.com/google/uuid"
)

type UserRepoForUser interface {
	Create(login string, payload dto.CreateUserDTO) (*model.User, error)
}

type UserService struct {
	repo UserRepoForUser
}

func NewUserService(repo UserRepoForUser) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) Create(payload dto.CreateUserDTO) (string, error) {
	login := uuid.New().String()

	_, err := s.repo.Create(login, payload)
	if err != nil {
		return "", err
	}

	return login, nil
}
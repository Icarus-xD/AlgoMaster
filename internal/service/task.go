package service

import (
	"encoding/json"
	"errors"

	"github.com/Icarus-xD/AlgoMaster/internal/dto"
	"github.com/Icarus-xD/AlgoMaster/internal/model"
)

const MAX_DEBT = 100000

type Emailer interface {
	SendEmail(emailType, to, subject string, data any) error
}

type UserRepoForTask interface {
	UpdateDebt(login string, price, maxDebt int,) (*model.User, error)
}

type TaskPriceRepo interface {
	GetPrice(t string) (*model.TaskPrice, error)
}

type TaskService struct {
	maxDebt int
	email Emailer
	userRepo UserRepoForTask
	priceRepo TaskPriceRepo
}

func NewTaskService(email Emailer, userRepo UserRepoForTask, priceRepo TaskPriceRepo) *TaskService {
	return &TaskService{
		maxDebt: MAX_DEBT,
		email: email,
		userRepo: userRepo,
		priceRepo: priceRepo,
	}
}

func (s *TaskService) Solve(payload dto.SolveTaskDTO) (any, error) {
	switch payload.Type {
	case "FMNUA":
		numbers, ok := payload.Data.([]int)
		if !ok {
			return nil, errors.New("wrong provided data, expected array of numbers")
		}

		price, err := s.getPrice(payload.Type)
		if err != nil {
			return nil, errors.New("failed to get price")
		}

		result := s.fmnua(numbers)

		user, err := s.pay(payload.Login, price, s.maxDebt)
		if err != nil {
			return nil, err
		}

		if user.Debt > s.maxDebt {
			_ = s.email.SendEmail("DebtExceeded", user.Email, "Debt Exceeded", struct{}{})
		}

		_ = s.email.SendEmail("TaskResult", user.Email, "Task Result", struct{
			Type string
			Input string
			Result string
		}{
			Type: payload.Type,
			Input: s.toString(numbers),
			Result: s.toString(result),
		})
		return result, nil
	default:
		return nil, errors.New("unknown type of task")
	}
}

func (s * TaskService) toString(result any) string {
	res, _ := json.Marshal(result)
	return string(res)
}	

func (s *TaskService) getPrice(t string) (int, error) {
	price, err := s.priceRepo.GetPrice(t)
	if err != nil {
		return 0, err
	}

	return price.Price, nil
}

func (s *TaskService) pay(login string, price, maxDebt int) (*model.User, error) {
	user, err := s.userRepo.UpdateDebt(login, price, maxDebt)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// Find Missing Numbers in Unsorted Array
func (s *TaskService) fmnua(numbers []int) []int {
	missingNumbers := []int{}

	flags := make([]bool, len(numbers))

	for _, num := range numbers {
		flags[num - 1] = true
	}

	for idx, flag := range flags {
		if !flag {
			missingNumbers = append(missingNumbers, idx + 1)
		}
	}

	return missingNumbers
}
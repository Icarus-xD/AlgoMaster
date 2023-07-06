package rest

import (
	"github.com/Icarus-xD/AlgoMaster/internal/dto"
	"github.com/gofiber/fiber/v2"
)

type errorResponse struct {
	Message string `json:"message"`
}

type UserService interface {
	Create(payload dto.CreateUserDTO) (string, error)
}

type TaskService interface {
	Solve(payload dto.SolveTaskDTO) (any, error)
}

type Handler struct {
	userService UserService
	taskService TaskService
}

func NewHandler(userService UserService, taskService TaskService) *Handler {
	return &Handler{
		userService: userService,
		taskService: taskService,
	}
}

func (h *Handler) DefineRoutes(app *fiber.App) {
	
	user := app.Group("/user")
	{
		user.Post("/create", h.createUser)
	}

	task := app.Group("/task")
	{
		task.Post("/solve", h.solveTask)
	}
}
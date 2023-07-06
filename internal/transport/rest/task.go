package rest

import (
	"encoding/json"
	"net/http"

	"github.com/Icarus-xD/AlgoMaster/internal/dto"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) solveTask(c *fiber.Ctx) error {
	var payload dto.SolveTaskDTO

	body := c.Body()
	err := json.Unmarshal(body, &payload)
	if err != nil {
		response := errorResponse{
			Message: err.Error(),
		}

		return c.Status(http.StatusBadRequest).JSON(response)
	}

	result, err := h.taskService.Solve(payload)
	if err != nil {
		response := errorResponse{
			Message: err.Error(),
		}

		return c.Status(http.StatusBadRequest).JSON(response)
	}

	return c.Status(http.StatusOK).JSON(map[string]any{
		"result": result,
	})
}
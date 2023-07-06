package rest

import (
	"encoding/json"
	"net/http"

	"github.com/Icarus-xD/AlgoMaster/internal/dto"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) createUser(c *fiber.Ctx) error {
	var payload dto.CreateUserDTO

	body := c.Body()
	err := json.Unmarshal(body, &payload)
	if err != nil {
		response := errorResponse{
			Message: err.Error(),
		}

		return c.Status(http.StatusBadRequest).JSON(response)
	}

	login, err := h.userService.Create(payload)
	if err != nil {
		response := errorResponse{
			Message: err.Error(),
		}

		return c.Status(http.StatusInternalServerError).JSON(response)
	}

	return c.Status(http.StatusCreated).JSON(map[string]string{
		"login": login,
	})
}
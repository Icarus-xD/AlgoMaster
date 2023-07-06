package dto

type CreateUserDTO struct {
	FirstName  string `json:"first_name"`
	MiddleName string `json:"middle_name"`
	LastName   string `json:"last_name"`
	Group      string `json:"group"`
	Email      string `json:"email"`
}
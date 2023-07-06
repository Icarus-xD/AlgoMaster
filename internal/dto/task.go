package dto

type SolveTaskDTO struct {
	Login string `json:"login"`
	Type  string `json:"type"`
	Data  any    `json:"data"`
}
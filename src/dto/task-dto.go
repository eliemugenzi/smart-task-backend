package dto

type Task struct {
	Title       string `json:"title" binding:"required"`
}
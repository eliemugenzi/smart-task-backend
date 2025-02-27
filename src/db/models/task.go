package models

import "fmt"

type Task struct {
	ID          uint   `gorm:"primaryKey,autoIncrement" json:"id"`
	Title       string `gorm:"not null" json:"title"`
	Description string `gorm:"not null" json:"description"`
	UserId      uint   `gorm:"not null" json:"user_id"`
	Status      string `gorm:"not null" json:"status"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	CreatedBy   User   `gorm:"foreignKey:UserId" json:"user"`
	Assignees   []User `gorm:"many2many:task_assignees" json:"assignees"`
}

func (this_ *Task) ToString() string {
	str := fmt.Sprintf("Task: \n{ \n\nID: %d,\nTitle: %s,\nDescription: %s,\nUserId: %d,\nStatus: %s,\nCreatedAt: %s, \nUpdatedAt: %s }", this_.ID, this_.Title, this_.Description, this_.UserId, this_.Status, this_.CreatedAt, this_.UpdatedAt)

	return str
}

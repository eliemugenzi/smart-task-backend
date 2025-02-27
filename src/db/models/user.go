package models

import (
	"fmt"
	"time"
)

type User struct {
	ID        uint      `gorm:"primaryKey,autoIncrement" json:"id"`
	FirstName string    `gorm:"not null" json:"first_name"`
	LastName  string    `gorm:"not null" json:"last_name"`
	Email     string    `gorm:"not null,unique" json:"email"`
	Password  string    `gorm:"not null" json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Tasks     []Task    `gorm:"foreignKey:UserId" json:"tasks"`
}

func (this_ *User) ToString() string {
	str := fmt.Sprintf("User: \n{ \n\nID: %d,\nFirstName: %s,\nLastName: %s,\nEmail: %s,\nCreatedAt: %s, \nUpdatedAt: %s }", this_.ID, this_.FirstName, this_.LastName, this_.Email, this_.Password, this_.CreatedAt, this_.UpdatedAt)

	return str
}

package main

import (
	"time"
)

type User struct {
	ID        int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Email     string    `json:"email" gorm:"unique;not null"`
	Password  string    `json:"password"`
	FirstName string    `json:"firstname"`
	LastName  string    `json:"lastname"`
	Phone     string    `json:"phone"`
	Birthday  time.Time `json:"birthday"`
	CreatedAt time.Time `json:"created_at"`
}

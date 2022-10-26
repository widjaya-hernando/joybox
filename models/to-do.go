package models

import "time"

type ToDo struct {
	ID          uint64    `json:"id"`
	UserID      uint64    `json:"user_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
	Status      bool      `json:"status"`
}

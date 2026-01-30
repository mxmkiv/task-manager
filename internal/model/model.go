package model

import "time"

type User struct {
	Id           int       `json:"id"`
	Login        string    `json:"login"`
	PasswordHash string    `json:"password_hash"`
	CreatedAt    time.Time `json:"created_at"`
	UpdateAt     time.Time `json:"update_at"`
}

type TaskStatus int

const (
	ToDoStatus TaskStatus = iota
	InProcessStatus
	DoneStatus
)

type Task struct {
	Id          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	OwnerId     int        `json:"owner_id"`
	Status      TaskStatus `json:"status"`
	Deadline    time.Time  `json:"deadline"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdateAt    time.Time  `json:"update_at"`
}

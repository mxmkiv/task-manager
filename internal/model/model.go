package model

import (
	"log"
	"task-manager/internal/auth"
	"time"
)

type User struct {
	Id           int       `json:"id"`
	Login        string    `json:"login"`
	PasswordHash string    `json:"-"`
	Role         string    `json:"role"`
	CreatedAt    time.Time `json:"created_at"`
	UpdateAt     time.Time `json:"update_at"`
}

type RoleType int

const (
	AdminType RoleType = iota
	UserType
)

func (r RoleType) RoleToString() string {
	mass := []string{"admin", "user"}
	return mass[r]
}

func NewUser(login, password string, encoder auth.HashEncoder) (*User, error) {

	hash, err := encoder.Encode(password)
	if err != nil {
		log.Println("generate hash error", err)
		return nil, err
	}
	return &User{
		Login:        login,
		PasswordHash: hash,
		Role:         AdminType.RoleToString(),
		CreatedAt:    time.Now(),
		UpdateAt:     time.Now(),
	}, nil
}

type TaskStatus int

const (
	ToDoStatus TaskStatus = iota
	InProcessStatus
	DoneStatus
)

func (t TaskStatus) StatusToString() string {
	mass := []string{"To do", "In process", "Done"}
	return mass[t]
}

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

// DTO
type RequestData struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type ResponseData struct {
	User  User   `json:"user"`
	Token string `json:"token"`
}

type UserData struct {
	Id        int       `json:"id"`
	Login     string    `json:"login"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"update_at"`
}

package model

import (
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id           int    `json:"id"`
	Login        string `json:"login"`
	PasswordHash string `json:"-"`
	//Role         string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"update_at"`
}

func NewUser(login, password string) (*User, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("generate hash error", err)
		return nil, err
	}
	return &User{
		Login:        login,
		PasswordHash: string(hash),
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

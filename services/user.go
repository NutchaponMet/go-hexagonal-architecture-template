package services

import "go-hexagonal-architecture/repository"

type UserResponse struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserService interface {
	CreateUser(repository.User) error
	GetUser(string) (*UserResponse, error)
}

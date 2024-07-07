package ports

import "go-hexagonal/core/domains"

type UserRepository interface {
	CreateUser(user *domains.User) error
	GetAll() ([]domains.User, error)
	GetByUserName(username string) (*domains.User, error)
}

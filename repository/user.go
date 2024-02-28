package repository

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique"`
	Password string `gorm:"unique"`
}

type UserRepository interface {
	CreateUser(User) (int64, error)
	GetByUsername(string) (User, error)
}

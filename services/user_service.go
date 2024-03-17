package services

import (
	"go-hexagonal-architecture/logs"
	"go-hexagonal-architecture/repository"

	"gorm.io/gorm"
)

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return userService{userRepo: userRepo}
}

func (s userService) CreateUser(user repository.User) error {
	err := s.userRepo.CreateUser(user)
	if err != nil {
		logs.Error(err)
		return err
	}
	return nil
}

func (s userService) GetUser(username string) (*UserResponse, error) {
	db_data, err := s.userRepo.GetByUsername(username)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			logs.Error(err)
			return nil, err
		}
		logs.Error(err)
		return nil, err
	}

	json_data := UserResponse{
		Username: db_data.Username,
		Password: db_data.Password,
	}

	return &json_data, nil
}

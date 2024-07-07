package services

import (
	"go-hexagonal/core/domains"
	"go-hexagonal/core/ports"
	"go-hexagonal/core/schemas"
	"go-hexagonal/pkg/errs"
	"go-hexagonal/pkg/logs"

	"gorm.io/gorm"
)

type userService struct {
	userRepo ports.UserRepository
}

func NewUserService(userRepo ports.UserRepository) ports.UserService {
	return userService{userRepo: userRepo}
}

func (s userService) NewUser(user *schemas.UserReq) error {
	newUser := domains.User{
		Username: user.Username,
		Password: user.Password,
		Email:    user.Email,
	}

	err := s.userRepo.CreateUser(&newUser)
	if err != nil {
		logs.Error(err)
		return errs.NewUnexpectedError("Unexpected Server Error")
	}
	return nil
}

func (s userService) GetUsers() ([]schemas.UserResp, error) {
	usersRepo, err := s.userRepo.GetAll()
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError("Unexpected Server Error")
	}

	if len(usersRepo) == 0 {
		logs.Error(gorm.ErrRecordNotFound)
		return nil, errs.NewNotfoundError("Record Not Found")
	}

	users := []schemas.UserResp{}
	for _, obj := range usersRepo {
		users = append(users, schemas.UserResp{
			Username: obj.Username,
			Password: obj.Password,
			Email:    obj.Email,
		})
	}
	return users, nil
}

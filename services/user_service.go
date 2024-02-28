package services

import "go-hexagonal-architecture/repository"

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) userService {
	return userService{userRepo: userRepo}
}

func (s userService) CreateUser(user repository.User) (int64, error) {
	result, err := s.userRepo.CreateUser(user)
	if err != nil {
		panic(err)
	}
	return result, nil
}

func (s userService) GetUser(username string) (UserResponse, error) {
	db_data, err := s.userRepo.GetByUsername(username)
	if err != nil {
		panic(err)
	}
	json_data := UserResponse{
		Username: db_data.Username,
		Password: db_data.Password,
	}
	return json_data, nil
}

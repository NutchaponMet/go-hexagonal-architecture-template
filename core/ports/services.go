package ports

import "go-hexagonal/core/schemas"

type UserService interface {
	NewUser(user *schemas.UserReq) error
	GetUsers() ([]schemas.UserResp, error)
}

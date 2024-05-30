package handlers

import (
	"go-hexagonal/repository"
	"go-hexagonal/services"

	"github.com/gofiber/fiber/v2"
)

type userHandler struct {
	userSrv services.UserService
}

func NewUserHandler(userSrv services.UserService) userHandler {
	return userHandler{userSrv: userSrv}
}

func (h userHandler) CreateUser(c *fiber.Ctx) error {
	userBody := repository.User{}
	err := c.BodyParser(&userBody)
	if err != nil {
		return err
	}
	err = h.userSrv.CreateUser(userBody)
	if err != nil {
		handleError(c, err)
	}
	return c.SendStatus(201)
}

func (h userHandler) GetUser(c *fiber.Ctx) error {
	user := c.Params("user")
	db_data, err := h.userSrv.GetUser(user)
	if err != nil {
		handleError(c, err)
	}
	return c.JSON(db_data)
}

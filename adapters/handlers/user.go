package handlers

import (
	"go-hexagonal/core/ports"
	"go-hexagonal/core/schemas"
	"go-hexagonal/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

type userHandler struct {
	userSrv ports.UserService
}

func NewUserHandler(userSrv ports.UserService) userHandler {
	return userHandler{userSrv: userSrv}
}

func (h userHandler) NewUser(c *fiber.Ctx) error {
	user := schemas.UserReq{}
	c.BodyParser(&user)

	err := h.userSrv.NewUser(&user)
	if err != nil {
		return utils.HandleError(c, err)
	}
	return c.SendStatus(201)
}

func (h userHandler) GetUsers(c *fiber.Ctx) error {
	respData, err := h.userSrv.GetUsers()
	if err != nil {
		return utils.HandleError(c, err)
	}
	return c.JSON(respData)
}

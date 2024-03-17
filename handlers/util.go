package handlers

import (
	"fmt"
	"go-hexagonal-architecture/errs"

	"github.com/gofiber/fiber/v2"
)

func handleError(c *fiber.Ctx, err error) {
	switch e := err.(type) {
	case errs.AppError:
		c.SendStatus(e.Code)
		fmt.Fprintln(c, e)
		// fmt.Println(e.Message)
	case error:
		c.SendStatus(fiber.StatusInternalServerError)
		fmt.Fprintln(c, e)
	}
}

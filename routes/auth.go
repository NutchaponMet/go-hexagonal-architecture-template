package routes

import (
	"go-hexagonal/handlers"
	"go-hexagonal/repository"
	"go-hexagonal/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Auth(db *gorm.DB) *fiber.App {
	app := fiber.New()

	userRepository := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userHandler := handlers.NewUserHandler(userService)

	app.Get("/user/:user?", userHandler.GetUser)
	app.Post("/signup", userHandler.CreateUser)
	return app
}

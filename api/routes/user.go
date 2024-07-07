package routes

import (
	"go-hexagonal/adapters/handlers"
	"go-hexagonal/adapters/repositories"
	"go-hexagonal/core/services"
	"go-hexagonal/pkg/db"

	"github.com/gofiber/fiber/v2"
)

func User(router fiber.Router) {

	userRepository := repositories.NewUserRepositoryDB(db.MARIADB)
	userService := services.NewUserService(userRepository)
	userHandler := handlers.NewUserHandler(userService)

	router.Get("/user", userHandler.GetUsers)
}

package main

import (
	"fmt"
	"go-hexagonal/logs"
	"go-hexagonal/middlewares"
	"go-hexagonal/routes"

	"github.com/gofiber/contrib/fiberzap/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	initTimeZone()
	initConfig()
	logs.LogInit()
	db = initDataBase()

}

func main() {

	app := fiber.New(fiber.Config{
		Prefork: viper.GetString("app.mode") == "release", // production Prefork : true
		// Prefork: false,
	})

	app.Use(middlewares.NewCorsOriginMiddleWare())

	if viper.GetString("app.mode") == "release" {
		app.Use(fiberzap.New(fiberzap.Config{Logger: logs.Log}))
	} else {
		app.Use(middlewares.NewLoggerMiddleWare())
	}

	app.Mount("/api/auth", routes.Auth(db))

	if viper.GetString("app.mode") == "release" {
		logs.Info("Fiber server start in production mode at port " + viper.GetString("app.port"))
		app.Listen(fmt.Sprintf(":%v", viper.GetInt("app.port")))
	} else {
		logs.Info("Fiber server start in debug mode at port " + viper.GetString("app.port"))
		app.Listen(fmt.Sprintf(":%v", viper.GetInt("app.port")))
	}
}

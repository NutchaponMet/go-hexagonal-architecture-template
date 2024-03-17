package main

import (
	"fmt"
	"go-hexagonal-architecture/logs"
	"go-hexagonal-architecture/middlewares"
	"go-hexagonal-architecture/routes"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
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

	app := fiber.New()

	app.Use(middlewares.NewCorsOriginMiddleWare())
	app.Use(middlewares.NewLoggerMiddleWare())

	app.Mount("/api/auth", routes.Auth(db))

	app.Listen(fmt.Sprintf(":%v", viper.GetInt("app.port")))
}

func initTimeZone() {
	ict, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		panic(err)
	}
	time.Local = ict
}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	viper.ReadInConfig()

}

func initDataBase() *gorm.DB {

	dsn := fmt.Sprintf(
		"host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Bangkok",
		viper.GetString("db.host"),
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.dbname"),
		viper.GetInt("db.port"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		DryRun:                 false, // ถ้ารค่าเป็น true จะแสดงเฉพาะ sql statment
		PrepareStmt:            true,
	})
	if err != nil {
		panic(err)
	}
	return db
}

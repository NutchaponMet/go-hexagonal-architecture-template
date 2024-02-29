package main

import (
	"fmt"
	"go-hexagonal-architecture/repository"
	"go-hexagonal-architecture/routes"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	initTimeZone()
	initConfig()

	app := fiber.New()

	db := initDataBase()

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
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
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
		DryRun: false, // ถ้ารค่าเป็น true จะแสดงเฉพาะ sql statment
	})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&repository.User{})
	return db
}

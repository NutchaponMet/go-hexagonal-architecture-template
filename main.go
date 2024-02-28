package main

import (
	"go-hexagonal-architecture/routes"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	app := fiber.New()

	db := initDataBase()

	app.Mount("/api/auth", routes.Auth(db))

	app.Listen(":8080")
}

func initDataBase() *gorm.DB {

	dsn := "host=localhost user=myuser password=mypassword dbname=mydb port=5432 sslmode=disable TimeZone=Asia/Bangkok"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		DryRun: true, // ถ้ารค่าเป็น true จะแสดงเฉพาะ sql statment
	})
	if err != nil {
		panic(err)
	}
	return db
}

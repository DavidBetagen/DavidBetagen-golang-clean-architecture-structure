package main

import (
	"golang-clean-architecture-structure/internal/infrastructure/applog"
	"golang-clean-architecture-structure/internal/infrastructure/config"
	"golang-clean-architecture-structure/internal/infrastructure/database"
	"golang-clean-architecture-structure/internal/route"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func init() {
	// ⚙️⚙️⚙️
	config.LoadConfig()
	// 📝📝📝
	applog.InitLog()
}

func main() {
	// ⚙️⚙️⚙️👩🏻‍💻
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOriginsFunc: func(origin string) bool {
			return config.AppConfig.ENVIRONMENT == "development"
		},
	}))

	// 📝📝📝 Init Logger
	file, err := os.OpenFile("log/logger.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer file.Close()
	app.Use(logger.New(logger.Config{
		Output: file,
	}))

	// 🛢️🛢️🛢️🛢️
	db, err := database.ConnectToDatabase()
	if err != nil {
		panic("Error connect to database. 🛢️💣💥")
	}

	// 🚦🚦🚦🚴
	route.Setup(app, db)

	// 🚀
	app.Listen(":" + config.AppConfig.PORT)
}

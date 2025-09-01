package main

import (
	"log"
	"student/database"
	"student/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {
	database.ConnectDB()
	if err := godotenv.Load(".env"); err != nil {
		panic("Error loading .env file")
	}

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET, POST, PUT, DELETE",
	}))

	app.Use(
		logger.New(logger.Config{
			Format:     "${time} ${status} - ${latency} ${method} ${path}\n",
			TimeFormat: "02-Jan-2006 15:04:05",
			TimeZone:   "Asia/Bangkok",
		}),
	)

	//CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET, POST, PUT, DELETE",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	router.SetUpRoutes(app)

	log.Fatal(app.Listen(":8080"))

}

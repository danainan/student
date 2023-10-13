package main

import (
	"log"
	"student/database"
	"student/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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

	router.SetUpRoutes(app)

	log.Fatal(app.Listen(":8000"))

}

package main

import (
	"github.com/daonham/go-fiber/controllers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization, Accept-Encoding, X-Requested-With",
		AllowMethods:     "GET, POST, PUT, DELETE, PATCH",
		AllowCredentials: true,
		ExposeHeaders:    "Content-Length",
		MaxAge:           86400,
	}))

	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestCompression,
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Post("/post", controllers.CreatePost)

	app.Listen(":8080")
}

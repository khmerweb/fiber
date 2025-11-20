// go mod init mainModule
// go get github.com/gofiber/fiber/v2
// go get -tool github.com/air-verse/air@latest

//package main

package handler

import (
	"fiber/admin"
	"fiber/frontend"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/template/html/v3"
	"github.com/joho/godotenv"
)

func Handler(w http.ResponseWriter, r *http.Request) {

	app := fiber.New()

	// Define your Fiber routes here
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello from Fiber on Vercel!")
	})

	adaptor.FiberApp(app)(w, r)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{Views: engine})
	app.Static("/", "./static")
	store := session.New()
	app.Use(cors.New())
	app.Use(func(c *fiber.Ctx) error {
		sess, err := store.Get(c)
		if err != nil {
			return err
		}

		sess.Set("name", "Guest")
		c.Locals("session", sess)

		return c.Next()
	})

	frontRoute := app.Group("/")
	adminRoute := app.Group("/admin")

	frontend.FrontRoutes(frontRoute)
	admin.AdminRoutes(adminRoute)

	app.Listen(":8000")
}

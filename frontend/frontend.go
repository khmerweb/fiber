package frontend

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func FrontRoutes(router fiber.Router) {
	router.Get("/", func(c *fiber.Ctx) error {
		sess := c.Locals("session").(*session.Session)
		name := sess.Get("name").(string)
		return c.Render("index", fiber.Map{
			"Title": "Homepage " + name,
		}, "base")
	})

}

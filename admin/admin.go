package admin

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func AdminRoutes(router fiber.Router) {

	router.Get("/", func(c *fiber.Ctx) error {
		sess := c.Locals("session").(*session.Session)
		name := sess.Get("name").(string)
		return c.Render("about", fiber.Map{
			"Title": "Adminpage " + name,
		}, "base")
	})

}

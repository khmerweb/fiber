package frontend

import (
	"fiber/db"

	"github.com/gofiber/fiber/v2"
	//"github.com/gofiber/fiber/v2/middleware/session"
)

func FrontRoutes(router fiber.Router) {
	router.Get("/", func(c *fiber.Ctx) error {
		//sess := c.Locals("session").(*session.Session)
		//name := sess.Get("name").(string)

		counts, posts := db.CountPosts()

		//fmt.Println(counts)
		return c.Render("index", fiber.Map{
			"Title":   "ដំណឹង​ល្អ ",
			"Counts":  counts,
			"Posts":   posts,
			"PageURL": "/",
		}, "base")
	})

}

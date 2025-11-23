package frontend

import (
	"encoding/json"
	"fiber/db"
	"fmt"

	"github.com/gofiber/fiber/v2"
	//"github.com/gofiber/fiber/v2/middleware/session"
)

func FrontRoutes(router fiber.Router) {
	router.Get("/", func(c *fiber.Ctx) error {
		//sess := c.Locals("session").(*session.Session)
		//name := sess.Get("name").(string)

		counts, posts := db.CountPosts()
		jsonDataString, err := json.Marshal(posts)
		if err != nil {
			fmt.Println("Error marshalling string data:", err)

		}
		playlists := string(jsonDataString)
		//fmt.Println(playlists)
		return c.Render("index", fiber.Map{
			"Title":     "ដំណឹង​ល្អ ",
			"Counts":    counts,
			"Posts":     posts,
			"Playlists": playlists,
			"PageURL":   "/",
		}, "base")
	})

}

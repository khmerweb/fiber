package frontend

import (
	"fiber/db"
	"fmt"

	"github.com/gofiber/fiber/v2"
	//"github.com/gofiber/fiber/v2/middleware/session"
)

func FrontRoutes(router fiber.Router) {
	router.Get("/", func(c *fiber.Ctx) error {
		//sess := c.Locals("session").(*session.Session)
		//name := sess.Get("name").(string)
		var count int64
		var counts []int64
		categories := []string{"news", "movie", "travel", "game", "sport", "doc", "food", "music", "distraction"}
		for _, value := range categories {
			count = db.CountPosts(value)
			counts = append(counts, count)
		}
		/*
			jsonDatas, err := json.Marshal(counts)
			if err != nil {
				fmt.Println("Error marshalling string data:", err)
				return err
			}

			jsonCounts := string(jsonDatas)
		*/
		fmt.Println(counts)
		return c.Render("index", fiber.Map{
			"Title":   "ដំណឹង​ល្អ ",
			"Counts":  counts,
			"PageURL": "/",
		}, "base")
	})

}

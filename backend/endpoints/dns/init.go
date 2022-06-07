package dns

import "github.com/gofiber/fiber/v2"

func Init(router fiber.Router) {
	router.All("/query", queryHandler)

	query := router.Group("/:channel_id/:channel_salt/query/", channel)
	query.Get("/", queryHandler)
	query.Post("/", queryHandler)
}

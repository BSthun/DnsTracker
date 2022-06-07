package record

import (
	"github.com/gofiber/fiber/v2"

	"backend/instances/models"
	"backend/types/responder"
)

func QueryHistoryHandler(c *fiber.Ctx) error {
	// * Parse channel
	channel := c.Locals("channel").(*models.Channel)

	// * Fetch last 1000 rows of log
	// log := channel.Logs[:1000]

	return c.JSON(responder.NewResponse(channel.Logs))
}

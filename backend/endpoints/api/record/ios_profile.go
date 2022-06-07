package record

import (
	"bytes"
	"fmt"

	"github.com/gofiber/fiber/v2"

	"backend/types/responder"
	"backend/utils/config"
	"backend/utils/constant"
)

func GenerateIosProfileHandler(c *fiber.Ctx) error {
	// * Parse query
	channelId := c.Query("id")
	channelSalt := c.Query("salt")

	var body bytes.Buffer
	if err := constant.IosProfileTemplate.Execute(&body, map[string]any{
		"SESSION_ID": channelId,
		"SERVER_URL": fmt.Sprintf("%s/dns/%s/%s/query", config.C.BackendUrl, channelId, channelSalt),
	}); err != nil {
		return &responder.GenericError{
			Message: "Unable to parse template",
		}
	}

	c.Set("Content-Type", "application/plist")
	c.Set("Content-Disposition", fmt.Sprintf("attachment; filename=dnstracker_session%s.mobileconfig", channelId))
	return c.Send(body.Bytes())
}

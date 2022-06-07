package middlewares

import (
	"strconv"

	"github.com/gofiber/fiber/v2"

	"backend/instances/hub"
	"backend/types/responder"
)

var Authen = func() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		// * Parse parameters
		chId := ctx.Get("X-CH-ID")
		if chId == "" {
			chId = ctx.Query("id")
		}

		channelId, err := strconv.ParseUint(chId, 10, 64)

		channelToken := ctx.Get("X-CH-TOKEN")
		if channelToken == "" {
			channelToken = ctx.Query("token")
		}

		if err != nil {
			return &responder.GenericError{
				Message: "Unable to parse Channel ID",
				Err:     err,
			}
		}

		// * Validate credential
		channel, ok := hub.Hub.Channels[channelId]
		if !ok || channel.Token != channelToken {
			return &responder.GenericError{
				Message: "Invalid channel credential",
			}
		}

		// / Passed validation

		// * Passing local variable
		ctx.Locals("channel", channel)

		return ctx.Next()
	}
}()

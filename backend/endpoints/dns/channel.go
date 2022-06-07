package dns

import (
	"strconv"

	"github.com/gofiber/fiber/v2"

	"backend/instances/hub"
	"backend/types/responder"
)

func channel(ctx *fiber.Ctx) error {
	// * Parse parameters
	channelId, err := strconv.ParseUint(ctx.Params("channel_id"), 10, 64)
	channelSalt := ctx.Params("channel_salt")

	if err != nil {
		return &responder.GenericError{
			Message: "Unable to parse Channel ID",
			Err:     err,
		}
	}

	// * Check parameters
	channel, ok := hub.Hub.Channels[channelId]
	if !ok {
		return &responder.GenericError{
			Message: "Channel ID does not exist",
		}
	}
	if channel.Salt != channelSalt {
		return &responder.GenericError{
			Message: "Channel Salt mismatched",
		}
	}

	// / Passed validation

	// * Passing local variable
	ctx.Locals("channel", channel)

	return ctx.Next()
}

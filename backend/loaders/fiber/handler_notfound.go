package fiber

import (
	"github.com/gofiber/fiber/v2"

	"backend/types/responder"
)

func notfoundHandler(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusNotFound).JSON(responder.ErrorResponse{
		Success: false,
		Message: "Endpoint Not Found",
		Code:    "NOT_FOUND",
	})
}

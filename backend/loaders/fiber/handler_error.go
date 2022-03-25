package fiber

import (
	"strings"

	"github.com/gofiber/fiber/v2"

	"backend/types/responder"
)

func errorHandler(ctx *fiber.Ctx, err error) error {
	// Case of *fiber.Error
	if e, ok := err.(*fiber.Error); ok {
		return ctx.Status(e.Code).JSON(responder.ErrorResponse{
			Success: false,
			Code:    strings.ReplaceAll(strings.ToUpper(e.Error()), " ", "_"),
			Message: e.Error(),
			Error:   e.Error(),
		})
	}

	// Case of *responder.GenericError
	if e, ok := err.(*responder.GenericError); ok {
		if e.Err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(responder.ErrorResponse{
				Success: false,
				Code:    e.Code,
				Message: e.Message,
				Error:   e.Err.Error(),
			})
		}

		return ctx.Status(fiber.StatusBadRequest).JSON(responder.ErrorResponse{
			Success: false,
			Code:    e.Code,
			Message: e.Message,
		})
	}

	// Unhandled error
	return ctx.Status(fiber.StatusInternalServerError).JSON(
		responder.ErrorResponse{
			Success: false,
			Code:    "SERVER_ERROR",
			Message: "Unknown server side error",
			Error:   err.Error(),
		},
	)
}

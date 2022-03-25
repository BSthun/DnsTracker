package middlewares

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"

	"backend/types/responder"
)

var limiterReached = &responder.GenericError{
	Message: "Rate limit reached, try again in a few minutes.",
}

var Limiter = func() fiber.Handler {
	config := limiter.Config{
		Next:       nil,
		Max:        6000,
		Expiration: 600 * time.Second,
		KeyGenerator: func(ctx *fiber.Ctx) string {
			return ctx.IP()
		},
		LimitReached: func(ctx *fiber.Ctx) error {
			return limiterReached
		},
	}

	return limiter.New(config)
}()

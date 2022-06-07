package middlewares

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

var Logger = func() fiber.Handler {
	config := logger.Config{
		TimeFormat: "02-Jan-2006 15:04:05",
		TimeZone:   "Asia/Bangkok",
		Output:     os.Stdout,
		Format:     "[${time}] ${method} ${path} [${queryParams}] [${body}] [${status}] [${resBody}]\n",
		Next: func(ctx *fiber.Ctx) bool {
			return true
		},
	}

	return logger.New(config)
}()

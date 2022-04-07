package fiber

import (
	"time"

	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"

	"backend/endpoints/api"
	"backend/endpoints/dns"
	"backend/loaders/fiber/middlewares"
	"backend/loaders/socketio"
	"backend/types/responder"
	"backend/utils/config"
	"backend/utils/logger"
)

var app *fiber.App

func Init() {
	// Initialize fiber instance
	app = fiber.New(fiber.Config{
		ErrorHandler:  errorHandler,
		Prefork:       false,
		StrictRouting: true,
		ServerHeader:  config.C.ServerHeader,
		ReadTimeout:   5 * time.Second,
		WriteTimeout:  5 * time.Second,
	})

	// Register root endpoint
	app.All("/", func(c *fiber.Ctx) error {
		return c.JSON(responder.InfoResponse{
			Success: true,
			Info:    "API_ROOT",
			Data:    nil,
		})
	})

	// Register API router
	apiGroup := app.Group("api/")

	// Apply middlewares to API router
	app.Use(middlewares.Cors)
	app.Use(middlewares.Recover)
	apiGroup.Use(middlewares.Limiter)

	// Apply endpoints to API router
	api.Init(apiGroup)

	// Register API router
	dnsGroup := app.Group("dns/")

	// Apply endpoints to API router
	dns.Init(dnsGroup)

	app.All("socket.io/", adaptor.HTTPHandlerFunc(socketio.Server.ServeHTTP))

	// Register not found handler
	app.Use(notfoundHandler)

	// Startup
	if err := app.Listen(config.C.Address); err != nil {
		logger.Log(logrus.Fatal, "LOAD FIBER FAILED: "+err.Error())
	}
}

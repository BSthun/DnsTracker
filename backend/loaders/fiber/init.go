package fiber

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"

	"backend/endpoints/api"
	"backend/endpoints/dns"
	"backend/loaders/fiber/middlewares"
	"backend/loaders/websocket"
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
		return c.JSON(responder.NewResponse("API_ROOT"))
	})

	// Register API router
	apiGroup := app.Group("api/")

	// Apply middlewares to API router
	app.Use(middlewares.Cors)
	app.Use(middlewares.Recover)
	app.Use(middlewares.Logger)
	apiGroup.Use(middlewares.Limiter)

	// Apply endpoints to API router
	api.Init(apiGroup)

	// Register DNS router
	dnsGroup := app.Group("dns/")

	// Apply endpoints to DNS router
	dns.Init(dnsGroup)

	// Init websocket
	websocketGroup := app.Group("ws/")

	// Apply endpoints to API router
	websocket.Init(websocketGroup)

	// Register not found handler
	app.Use(notfoundHandler)

	// Startup
	if err := app.Listen(config.C.Address); err != nil {
		logger.Log(logrus.Fatal, "LOAD FIBER FAILED: "+err.Error())
	}
}

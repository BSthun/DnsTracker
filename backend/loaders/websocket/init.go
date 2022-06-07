package websocket

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/sirupsen/logrus"

	"backend/instances/hub"
	"backend/utils/logger"
)

func Init(router fiber.Router) {
	router.Use("/", func(c *fiber.Ctx) error {
		// IsWebSocketUpgrade returns true if the client
		// requested upgrade to the WebSocket protocol.
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	router.Get("/log", websocket.New(func(conn *websocket.Conn) {
		// websocket.Conn bindings https://pkg.go.dev/github.com/fasthttp/websocket?tab=doc#pkg-index
		hub.Hub.Serve(conn)
	}))

	logger.Log(logrus.Info, "LOADED WEBSOCKET SERVER")
}

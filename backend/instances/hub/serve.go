package hub

import (
	"strconv"

	"github.com/gofiber/websocket/v2"
	"github.com/sirupsen/logrus"

	"backend/utils/logger"
)

func (r *hub) Serve(conn *websocket.Conn) {
	// Defer connection
	defer func() {
		// Close connection
		_ = conn.Close()
	}()

	// * Check connection credentials
	channelId, err := strconv.ParseUint(conn.Query("channel_id"), 10, 64)
	if err != nil {
		return
	}

	channel, ok := r.Channels[channelId]
	if !ok || channel.Token != conn.Query("channel_token") {
		return
	}

	// * Connection switch
	if channel.Conn != nil && channel.Conn.Conn != nil {
		if err := channel.Conn.Close(); err != nil {
			logger.Log(logrus.Warn, "UNHANDLED CONN SWITCH: "+err.Error())
		}
	}
	channel.Conn = conn

	// * Reading message loop
	for {
		t, _, err := conn.ReadMessage()
		if err != nil {
			break
		}
		if t != websocket.TextMessage {
			continue
		}
	}
}

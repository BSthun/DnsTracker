package account

import (
	"fmt"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"

	"backend/instances/hub"
	"backend/instances/models"
	"backend/types/body"
	"backend/types/responder"
	"backend/utils/config"
	"backend/utils/text"
)

func NewSessionHandler(c *fiber.Ctx) error {
	channel := &models.Channel{
		Id:        hub.Hub.GetIncrementValue(),
		Salt:      *text.GenerateString(text.GenerateStringSet.MixedAlphaNum, 6),
		Token:     *text.GenerateString(text.GenerateStringSet.MixedAlphaNum, 32),
		Logs:      []*models.Log{},
		Conn:      nil,
		ConnMutex: &sync.Mutex{},
	}

	hub.Hub.Channels[channel.Id] = channel

	time.Sleep(2 * time.Second)

	return c.JSON(responder.NewResponse(&body.SessionResponse{
		ChannelId:    channel.Id,
		ChannelToken: channel.Token,
		Salt:         channel.Salt,
		BaseUrl:      config.C.BackendUrl,
		WebsocketUrl: config.C.WebSocketUrl,
		DohUrl:       fmt.Sprintf("%s/dns/%d/%s/query", config.C.BackendUrl, channel.Id, channel.Salt),
	}))
}

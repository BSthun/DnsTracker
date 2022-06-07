package models

import (
	"fmt"
	"sync"

	"github.com/gofiber/websocket/v2"
	"github.com/sirupsen/logrus"

	websocketLoader "backend/loaders/websocket/messages"
	"backend/utils/logger"
)

type Channel struct {
	Id        uint64          `json:"id"`
	Salt      string          `json:"salt"`
	Token     string          `json:"token"`
	Logs      []*Log          `json:"logs"`
	Conn      *websocket.Conn `json:"-"`
	ConnMutex *sync.Mutex     `json:"-"`
}

func (r *Channel) Emit(payload *websocketLoader.OutboundMessage) {
	if r.Conn == nil || r.Conn.Conn == nil {
		return
	}
	r.ConnMutex.Lock()
	if err := r.Conn.WriteJSON(payload); err != nil {
		logger.Log(logrus.Warn, fmt.Sprintf("WRITING MESSAGE FAILURE FOR PLAYER %s: %s", r.Id, err.Error()))
	}
	r.ConnMutex.Unlock()
}

package dns

import (
	"time"

	"backend/instances/models"
	"backend/loaders/dns"
	"backend/loaders/websocket/messages"
)

func log(channel *models.Channel, req *dns.Record) {
	log := &models.Log{
		Hostname: req.Request.Question[0].Name,
		Status:   req.Response.Rcode,
		Time:     time.Now(),
	}

	if len(channel.Logs) > 1000000 {
		channel.Logs = channel.Logs[1000:]
	}

	channel.Emit(&messages.OutboundMessage{
		Event:   messages.LogUpdate,
		Payload: log,
	})

	channel.Logs = append(channel.Logs, log)
}

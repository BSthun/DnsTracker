package dns

import (
	"context"
	"math/rand"

	"github.com/gofiber/fiber/v2"

	"backend/utils/config"
)

func Query(record *Record) *Error {
	var err error
	ctx := context.Background()
	numUpstreams := len(config.C.DnsUpstreams)

	for i := int64(0); i < config.C.DnsResolveTires; i++ {
		upstream := config.C.DnsUpstreams[rand.Intn(numUpstreams)]
		record.CurrentUpstream = upstream.Address

		switch upstream.Proto {
		default:
			return &Error{
				Code:    fiber.StatusInternalServerError,
				Message: "Invalid upstream protocol type",
			}
		case "tcp":
			record.Response, _, err = TcpClient.ExchangeContext(ctx, record.Request, upstream.Address)
		case "udp":
			record.Response, _, err = UdpClient.ExchangeContext(ctx, record.Request, upstream.Address)
			// TODO: IX-FR request and TCP fallback
		}

		if err == nil {
			return nil
		}
	}
	return &Error{
		Code:    fiber.StatusInternalServerError,
		Message: "No upstream response",
	}
}

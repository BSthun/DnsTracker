package dns

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"

	"backend/instances/models"
	"backend/loaders/dns"
)

func queryHandler(ctx *fiber.Ctx) error {
	// Retrieve content type
	contentType := ctx.Get("Content-Type")
	if ct := ctx.FormValue("ct"); ct != "" {
		contentType = ct
	}

	if contentType == "" {
		if ctx.FormValue("name") != "" {
			contentType = "application/dns-json"
		} else if ctx.FormValue("dns") != "" {
			contentType = "application/dns-message"
		}
	}

	// Retrieve response type
	var responseType string
	for _, responseCandidate := range strings.Split(ctx.Get("Accept"), ",") {
		responseCandidate = strings.SplitN(responseCandidate, ";", 2)[0]
		if responseCandidate == "application/json" {
			responseType = "application/json"
			break
		} else if responseCandidate == "application/dns-udpwireformat" {
			responseType = "application/dns-message"
			break
		} else if responseCandidate == "application/dns-message" {
			responseType = "application/dns-message"
			break
		}
	}
	if responseType == "" {
		if contentType == "application/dns-json" {
			responseType = "application/json"
		} else if contentType == "application/dns-message" {
			responseType = "application/dns-message"
		} else if contentType == "application/dns-udpwireformat" {
			responseType = "application/dns-message"
		}
	}

	// Create DNS query struct
	req := new(dns.Record)
	err := new(dns.Error)
	if contentType == "application/dns-json" {
		err = dns.GoogleRequestFormatter(req, ctx)
	} else if contentType == "application/dns-message" || contentType == "application/dns-udpwireformat" {
		err = dns.IetfRequestFormatter(req, ctx)
	} else {
		return ctx.Status(fiber.StatusUnsupportedMediaType).SendString(fmt.Sprintf("Invalid argument value: \"ct\" = %q", contentType))
	}

	err = dns.Query(req)
	if err != nil {
		return ctx.Status(err.Code).SendString(err.Message)
	}

	// Log dns answer
	if ctx.Locals("channel") != nil {
		log(ctx.Locals("channel").(*models.Channel), req)
	}

	// Response dns answer
	if responseType == "application/json" {
		return dns.GoogleResponseFormatter(req, ctx)
	} else if responseType == "application/dns-message" {
		return dns.IetfResponseFormatter(req, ctx)
	}
	return ctx.Status(fiber.StatusInternalServerError).SendString("Unknown response Content-Type")
}

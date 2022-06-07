package dns

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	jsondns "github.com/m13253/dns-over-https/v2/json-dns"
	"github.com/miekg/dns"
)

// IETF-style DNS formatter
// Refactored from https://github.com/m13253/dns-over-https/blob/master/doh-server/ietf.go

func IetfRequestFormatter(record *Record, ctx *fiber.Ctx) *Error {
	requestBase64 := ctx.FormValue("dns")
	requestBinary, err := base64.RawURLEncoding.DecodeString(requestBase64)
	if err != nil {
		return &Error{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("Invalid argument value: \"dns\" = %q", requestBase64),
		}
	}
	if len(requestBinary) == 0 && (ctx.Get("Content-Type") == "application/dns-message" || ctx.Get("Content-Type") == "application/dns-udpwireformat") {
		requestBinary = ctx.Body()
	}
	if len(requestBinary) == 0 {
		return &Error{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("Invalid argument value: \"dns\""),
		}
	}

	// TODO: patchDNSCryptProxyReqID

	msg := new(dns.Msg)
	err = msg.Unpack(requestBinary)
	if err != nil {
		return &Error{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("DNS packet parse failure (%s)", err.Error()),
		}
	}

	transactionID := msg.Id
	msg.Id = dns.Id()
	opt := msg.IsEdns0()
	if opt == nil {
		opt = new(dns.OPT)
		opt.Hdr.Name = "."
		opt.Hdr.Rrtype = dns.TypeOPT
		opt.SetUDPSize(dns.DefaultMsgSize)
		opt.SetDo(false)
		msg.Extra = append([]dns.RR{opt}, msg.Extra...)
	}
	var edns0Subnet *dns.EDNS0_SUBNET
	for _, option := range opt.Option {
		if option.Option() == dns.EDNS0SUBNET {
			edns0Subnet = option.(*dns.EDNS0_SUBNET)
			break
		}
	}
	isTailored := edns0Subnet == nil

	if edns0Subnet == nil {
		ednsClientFamily := uint16(0)
		ednsClientAddress := findClientIP(ctx)
		ednsClientNetmask := uint8(255)
		if ednsClientAddress != nil {
			if ipv4 := ednsClientAddress.To4(); ipv4 != nil {
				ednsClientFamily = 1
				ednsClientAddress = ipv4
				ednsClientNetmask = 24
				// TODO: Use precise IP config
			} else {
				ednsClientFamily = 2
				ednsClientNetmask = 128
				// TODO: Use precise IP config
			}
			edns0Subnet = new(dns.EDNS0_SUBNET)
			edns0Subnet.Code = dns.EDNS0SUBNET
			edns0Subnet.Family = ednsClientFamily
			edns0Subnet.SourceNetmask = ednsClientNetmask
			edns0Subnet.SourceScope = 0
			edns0Subnet.Address = ednsClientAddress
			opt.Option = append(opt.Option, edns0Subnet)
		}
	}

	record.Request = msg
	record.TransactionID = transactionID
	record.IsTailored = isTailored
	return nil
}

func IetfResponseFormatter(req *Record, ctx *fiber.Ctx) error {
	respJSON := jsondns.Marshal(req.Response)
	req.Response.Id = req.TransactionID
	respBytes, err := req.Response.Pack()
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).SendString(fmt.Sprintf("DNS packet construct failure (%s)", err.Error()))
	}

	ctx.Set("Content-Type", "application/dns-message")
	now := time.Now().UTC().Format(http.TimeFormat)
	ctx.Set("Date", now)
	ctx.Set("Last-Modified", now)
	ctx.Set("Vary", "Accept")

	// TODO: Patch firefox content type

	if respJSON.HaveTTL {
		if req.IsTailored {
			ctx.Set("Cache-Control", "private, max-age="+strconv.FormatUint(uint64(respJSON.LeastTTL), 10))
		} else {
			ctx.Set("Cache-Control", "public, max-age="+strconv.FormatUint(uint64(respJSON.LeastTTL), 10))
		}
		ctx.Set("Expires", respJSON.EarliestExpires.Format(http.TimeFormat))
	}

	if respJSON.Status == dns.RcodeServerFailure {
		return ctx.Status(http.StatusServiceUnavailable).SendString(fmt.Sprintf("received server failure from upstream %s: %v\n", req.CurrentUpstream, req.Response))
	}

	_, err = ctx.Write(respBytes)
	return err
}

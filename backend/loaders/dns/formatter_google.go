package dns

import (
	"fmt"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	jsondns "github.com/m13253/dns-over-https/v2/json-dns"
	"github.com/miekg/dns"
	"golang.org/x/net/idna"
)

// Google-style DNS formatter
// Refactored from https://github.com/m13253/dns-over-https/blob/master/doh-server/google.go

func GoogleRequestFormatter(record *Record, ctx *fiber.Ctx) *Error {
	// Parse query name
	name, err := parseQueryName(ctx.FormValue("name"))
	if err != nil {
		return err
	}

	// Parse query type
	typ, err := parseQueryType(ctx.FormValue("type"))
	if err != nil {
		return err
	}

	// Parse query cd
	cd, err := parseQueryCd(ctx.FormValue("cd"))
	if err != nil {
		return err
	}

	// Parse query EDNS subnet
	edns, err := parseQueryEdnsSubnet(ctx.FormValue("edns_client_subnet"), ctx)
	if err != nil {
		return err
	}

	// Create DNS option struct
	opt := new(dns.OPT)
	opt.Hdr.Name = "."
	opt.Hdr.Rrtype = dns.TypeOPT
	opt.SetUDPSize(dns.DefaultMsgSize)
	opt.SetDo(true)
	if edns != nil {
		opt.Option = append(opt.Option, edns)
	}

	// Create DNS message
	msg := new(dns.Msg)
	msg.SetQuestion(dns.Fqdn(name), typ)
	msg.CheckingDisabled = cd
	msg.Extra = append(msg.Extra, opt)

	record.Request = msg
	record.IsTailored = edns == nil
	return nil
}

func GoogleResponseFormatter(req *Record, ctx *fiber.Ctx) error {
	respJSON := jsondns.Marshal(req.Response)
	now := time.Now().UTC().Format(http.TimeFormat)

	ctx.Set("Content-Type", "application/json; charset=UTF-8")
	ctx.Set("Date", now)
	ctx.Set("Last-Modified", now)
	ctx.Set("Vary", "Accept")

	if respJSON.HaveTTL {
		if req.IsTailored {
			ctx.Set("Cache-Control", "private, max-age="+strconv.FormatUint(uint64(respJSON.LeastTTL), 10))
		} else {
			ctx.Set("Cache-Control", "public, max-age="+strconv.FormatUint(uint64(respJSON.LeastTTL), 10))
		}
		ctx.Set("Expires", respJSON.EarliestExpires.Format(http.TimeFormat))
	}

	if respJSON.Status == dns.RcodeServerFailure {
		ctx.Status(http.StatusServiceUnavailable)
	}

	return ctx.JSON(respJSON)
}

func parseQueryName(name string) (string, *Error) {
	if name == "" {
		return "", &Error{
			Code:    http.StatusBadRequest,
			Message: "No value for query name",
		}
	}

	if punycode, err := idna.ToASCII(name); err != nil {
		return "", &Error{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("Invalid argument value: \"name\" = %q (%s)", name, err.Error()),
		}
	} else {
		name = punycode
	}

	return name, nil
}

func parseQueryType(typStr string) (uint16, *Error) {
	var typ uint16
	if typStr == "" {
		typ = 1
	} else if v, err := strconv.ParseUint(typStr, 10, 16); err == nil {
		typ = uint16(v)
	} else if v, ok := dns.StringToType[strings.ToUpper(typStr)]; ok {
		typ = v
	} else {
		return typ, &Error{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("Invalid argument value: \"type\" = %q", typStr),
		}
	}
	return typ, nil
}

func parseQueryCd(cdStr string) (bool, *Error) {
	cd := false
	if cdStr == "1" || strings.EqualFold(cdStr, "true") {
		cd = true
	}
	if cdStr != "0" && cdStr != "" && !strings.EqualFold(cdStr, "false") {
		return cd, &Error{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("Invalid argument value: \"cd\" = %q", cdStr),
		}
	}
	return cd, nil
}

func parseQueryEdnsSubnet(ednsClientSubnet string, ctx *fiber.Ctx) (*dns.EDNS0_SUBNET, *Error) {
	ednsClientFamily := uint16(0)
	ednsClientAddress := net.IP(nil)
	ednsClientNetmask := uint8(255)
	if ednsClientSubnet != "" {
		if ednsClientSubnet == "0/0" {
			ednsClientSubnet = "0.0.0.0/0"
		}
		slash := strings.IndexByte(ednsClientSubnet, '/')
		if slash < 0 {
			ednsClientAddress = net.ParseIP(ednsClientSubnet)
			if ednsClientAddress == nil {
				return nil, &Error{
					Code:    http.StatusBadRequest,
					Message: fmt.Sprintf("Invalid argument value: \"edns_client_subnet\" = %q", ednsClientSubnet),
				}
			}
			if ipv4 := ednsClientAddress.To4(); ipv4 != nil {
				ednsClientFamily = 1
				ednsClientAddress = ipv4
				ednsClientNetmask = 24
			} else {
				ednsClientFamily = 2
				ednsClientNetmask = 56
			}
		} else {
			ednsClientAddress = net.ParseIP(ednsClientSubnet[:slash])
			if ednsClientAddress == nil {
				return nil, &Error{
					Code:    http.StatusBadRequest,
					Message: fmt.Sprintf("Invalid argument value: \"edns_client_subnet\" = %q", ednsClientSubnet),
				}
			}
			if ipv4 := ednsClientAddress.To4(); ipv4 != nil {
				ednsClientFamily = 1
				ednsClientAddress = ipv4
			} else {
				ednsClientFamily = 2
			}
			netmask, err := strconv.ParseUint(ednsClientSubnet[slash+1:], 10, 8)
			if err != nil {
				return nil, &Error{
					Code:    http.StatusBadRequest,
					Message: fmt.Sprintf("Invalid argument value: \"edns_client_subnet\" = %q", ednsClientSubnet),
				}
			}
			ednsClientNetmask = uint8(netmask)
		}
	} else {
		ednsClientAddress = findClientIP(ctx)
		if ednsClientAddress == nil {
			ednsClientNetmask = 0
		} else if ipv4 := ednsClientAddress.To4(); ipv4 != nil {
			ednsClientFamily = 1
			ednsClientAddress = ipv4
			ednsClientNetmask = 24
		} else {
			ednsClientFamily = 2
			ednsClientNetmask = 56
		}
	}

	if ednsClientAddress == nil {
		return nil, nil
	}

	edns0Subnet := new(dns.EDNS0_SUBNET)
	edns0Subnet.Code = dns.EDNS0SUBNET
	edns0Subnet.Family = ednsClientFamily
	edns0Subnet.SourceNetmask = ednsClientNetmask
	edns0Subnet.SourceScope = 0
	edns0Subnet.Address = ednsClientAddress
	return edns0Subnet, nil
}

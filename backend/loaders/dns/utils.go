package dns

import (
	"net"
	"strings"

	"github.com/gofiber/fiber/v2"
	jsondns "github.com/m13253/dns-over-https/v2/json-dns"

	"backend/utils/config"
)

func findClientIP(ctx *fiber.Ctx) net.IP {
	XForwardedFor := ctx.Get("X-Forwarded-For")
	if XForwardedFor != "" {
		for _, addr := range strings.Split(XForwardedFor, ",") {
			addr = strings.TrimSpace(addr)
			ip := net.ParseIP(addr)
			if jsondns.IsGlobalIP(ip) {
				return ip
			}
		}
	}
	XRealIP := ctx.Get("X-Real-IP")
	if XRealIP != "" {
		addr := strings.TrimSpace(XRealIP)
		ip := net.ParseIP(addr)
		if config.C.EcsAllowNonGlobalIp || jsondns.IsGlobalIP(ip) {
			return ip
		}
	}

	remoteAddr, err := net.ResolveTCPAddr("tcp", ctx.IP())
	if err != nil {
		return nil
	}
	ip := remoteAddr.IP
	if config.C.EcsAllowNonGlobalIp || jsondns.IsGlobalIP(ip) {
		return ip
	}
	return nil
}

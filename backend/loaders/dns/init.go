package dns

import (
	"time"

	"github.com/miekg/dns"
	"github.com/sirupsen/logrus"

	"backend/utils/config"
	"backend/utils/logger"
)

var UdpClient *dns.Client
var TcpClient *dns.Client

func Init() {
	timeout := time.Duration(config.C.DnsResolveTimeout) * time.Second

	UdpClient = &dns.Client{
		Net:     "udp",
		UDPSize: dns.DefaultMsgSize,
		Timeout: timeout,
	}

	TcpClient = &dns.Client{
		Net:     "udp",
		UDPSize: dns.DefaultMsgSize,
		Timeout: timeout,
	}

	logger.Log(logrus.Info, "LOADED DNS CLIENT SERVICE")
}

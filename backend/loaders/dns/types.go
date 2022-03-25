package dns

import "github.com/miekg/dns"

type Record struct {
	request         *dns.Msg
	response        *dns.Msg
	transactionID   uint16
	currentUpstream string
	isTailored      bool
}

type Error struct {
	Code    int
	Message string
}

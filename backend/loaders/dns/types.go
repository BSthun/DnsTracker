package dns

import "github.com/miekg/dns"

type Record struct {
	Request         *dns.Msg
	Response        *dns.Msg
	TransactionID   uint16
	CurrentUpstream string
	IsTailored      bool
}

type Error struct {
	Code    int
	Message string
}

package config

import "backend/types/enum"

type config struct {
	Environment enum.Environment
	LogLevel    uint32

	Address         string
	SocketioAddress string
	ServerHeader    string
	Cors            []string

	JwtSecret string

	RedisAddress  string
	RedisPassword string
	RedisDb       int

	DnsAuthority      string
	DnsResolveTimeout int64
	DnsResolveTires   int64
	DnsUpstreams      []struct {
		Proto   string
		Address string
	}

	EcsAllowNonGlobalIp bool
}

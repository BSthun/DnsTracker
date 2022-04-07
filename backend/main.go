package main

import (
	"math/rand"
	"time"

	"backend/loaders/dns"
	"backend/loaders/fiber"
	"backend/loaders/socketio"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// TODO: Store query on Redis
	// redis.Init()

	dns.Init()
	socketio.Init()
	fiber.Init()
}

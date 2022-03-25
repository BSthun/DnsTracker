package main

import (
	"math/rand"
	"time"

	"backend/loaders/dns"
	"backend/loaders/fiber"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// TODO: Store query on Redis
	// redis.Init()

	dns.Init()
	fiber.Init()
}

package main

import (
	"flag"
	"log"

	"github.com/therusetiawan/http-rate-limiter/internal/api"
	"github.com/therusetiawan/http-rate-limiter/pkg/redis"
)

var (
	listenPort = flag.String("listen-port", "9000", "Port where app listen to")
	redisUrl   = flag.String("redis-url", "localhost:6379", "Connection string to redis")
	rateLimit  = flag.Int64("rate-limit", 5, "Max request per second")
)

func main() {
	flag.Parse()

	// redis connection
	redisConfig := redis.Config{
		Addr: *redisUrl,
	}
	err := redisConfig.NewConnection()
	if err != nil {
		log.Printf("redis connection failed. rate limiter will not work")
	}

	// server configurations
	apiConfig := api.Config{
		ListenPort: *listenPort,
		RateLimit:  *rateLimit,
	}

	apiConfig.Start() // start server
}

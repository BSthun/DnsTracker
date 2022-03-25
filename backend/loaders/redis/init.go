package redis

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"

	"backend/utils/config"
	"backend/utils/logger"
)

var Client *redis.Client

func Init() {
	client := redis.NewClient(&redis.Options{
		Addr:     config.C.RedisAddress,
		Password: config.C.RedisPassword,
		DB:       config.C.RedisDb,
	})

	ctx := context.Background()
	if status := client.Set(ctx, "key", "value", 0); status.Err() != nil {
		logger.Log(logrus.Fatal, "LOAD REDIS FAILED: "+status.Err().Error())
	} else {
		logger.Log(logrus.Info, "INITIALIZED REDIS CONNECTION")
		Client = client
	}
}

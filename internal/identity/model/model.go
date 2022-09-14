package model

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v9"
	"github.com/sirupsen/logrus"

	"github.com/jdxj/cyber-wagon/config"
)

var (
	redisClient *redis.Client
)

func Init(cfg config.Redis) {
	rc := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		DB:   cfg.DB,
	})

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := rc.Ping(ctx).Err(); err != nil {
		logrus.Fatalln(err)
	}
	redisClient = rc
}

package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/bwmarrin/snowflake"
	"github.com/go-redis/redis/v9"
	uuid "github.com/satori/go.uuid"

	"github.com/jdxj/cyber-wagon/config"
)

func TestUUID(t *testing.T) {
	u := uuid.NewV4()
	fmt.Printf("u: %s\n", u)
	fmt.Printf("%d\n", len(u.Bytes()))
}

func TestSnowflake(t *testing.T) {
	n, err := snowflake.NewNode(1)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("%d\n", n.Generate())
	fmt.Printf("%d\n", n.Generate())
}

func TestPingRedis(t *testing.T) {
	config.Init("../config/test.yaml")
	cfg := config.GetRedis()
	rc := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		DB:   cfg.DB,
	})
	cmd := rc.Ping(context.Background())
	result, err := cmd.Result()
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("result: %+v\n", result)

	intCmd := rc.HIncrBy(context.Background(), "increment_id", "1", 10)
	id, err := intCmd.Result()
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("id: %d\n", id)
}

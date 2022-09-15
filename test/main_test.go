package test

import (
	"context"
	"crypto/rand"
	"crypto/sha512"
	"encoding/base32"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/bwmarrin/snowflake"
	"github.com/go-redis/redis/v9"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/sha3"

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

func TestSha(t *testing.T) {
	sha512.New()
	sha3.New512()
}

func TestFile(t *testing.T) {
	filepath.Base("")
	fmt.Printf("temp: %s\n", os.TempDir())
	fi, err := os.Stat("abc/LPXVBJ6JYZY3FTL5")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("not exist")
		}
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("%+v\n", fi)
}

func TestBase64(t *testing.T) {
	buf := make([]byte, 10)
	_, err := rand.Read(buf)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	s := base32.StdEncoding.EncodeToString(buf)
	fmt.Printf("s: %s\n", s)
}

func TestMkdir(t *testing.T) {
	err := os.Mkdir("test-create-dir", os.ModePerm)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
}

func TestRename(t *testing.T) {
	err := os.Rename("abc", "test-create-dir/def")
	if err != nil {
		t.Fatalf("%s\n", err)
	}
}

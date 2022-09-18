package test

import (
	"context"
	"crypto/rand"
	"crypto/sha512"
	"encoding/base32"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"testing"
	"time"

	"github.com/bwmarrin/snowflake"
	"github.com/go-redis/redis/v9"
	amqp "github.com/rabbitmq/amqp091-go"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/sha3"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"

	"github.com/jdxj/cyber-wagon/config"
	gateway "github.com/jdxj/cyber-wagon/internal/gateway/proto"
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

func TestFmt(t *testing.T) {
}

func TestRabbitMQ(t *testing.T) {
	conn, err := amqp.Dial("amqp://guest:guest@192.168.50.200:5672/")
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	defer ch.Close()

	err = ch.ExchangeDeclare("test-ex", "topic", false, true, false, false, nil)
	if err != nil {
		t.Fatalf("%s\n", err)
	}

	queue, err := ch.QueueDeclare("hello", false, true, false, false, nil)
	if err != nil {
		t.Fatalf("%s\n", err)
	}

	err = ch.QueueBind(queue.Name, "hello_bk", "test-ex", false, nil)
	if err != nil {
		t.Fatalf("%s\n", err)
	}

	go func() {
		for {
			err := ch.PublishWithContext(context.Background(), "test-ex", "hello_bk", false,
				false, amqp.Publishing{Body: []byte("def")})
			if err != nil {
				fmt.Printf("publish err: %s\n", err)
			}
			time.Sleep(time.Second)
		}
	}()

	deliveryCh, err := ch.Consume(queue.Name, "", false, false, false, false, nil)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	go func() {
		for d := range deliveryCh {
			fmt.Printf("body: %s\n", d.Body)
			err := d.Ack(false)
			if err != nil {
				fmt.Printf("ack err: %s", err)
			}
		}
	}()

	time.Sleep(time.Hour)
}

func TestProto(t *testing.T) {
	msg := &gateway.ClientMsg{}
	fmt.Printf("%#v\n", msg.ProtoReflect().Descriptor().FullName())
	fmt.Printf("%#v\n", msg.ProtoReflect().Descriptor().Name())

	data, err := proto.Marshal(msg)
	if err != nil {
		t.Fatalf("%s\n", err)
	}

	a := &anypb.Any{
		TypeUrl: string(msg.ProtoReflect().Descriptor().FullName()),
		Value:   data,
	}
	data, err = proto.Marshal(a)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
}

func TestPC(t *testing.T) {
	pc := make([]uintptr, 6)
	n := runtime.Callers(0, pc)
	f := runtime.CallersFrames(pc[:n])
	for frame, ok := f.Next(); ok; frame, ok = f.Next() {
		t.Logf("file: %s, line: %d\n", frame.File, frame.Line)
	}
}

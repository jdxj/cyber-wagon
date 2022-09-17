package util

import (
	"context"
	"fmt"
	"testing"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/jdxj/cyber-wagon/config"
)

func TestCreateExchange(t *testing.T) {
	InitMQ(config.Rabbitmq{
		Host: "192.168.50.200",
		Port: 5672,
		User: "guest",
		Pass: "guest",
	})

	ch, err := MQ.Channel()
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	defer ch.Close()

	go func() {
		for {
			err = publish(context.Background(), ch, LogicBindingKey, amqp.Publishing{
				Body: []byte("abc"),
			})
			if err != nil {
				fmt.Printf("publish err: %s", err)
			}

			time.Sleep(time.Second)
		}
	}()

	dc, err := consume(ch, LogicQueue, "def")
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	go func() {
		for d := range dc {
			fmt.Printf("msg: %s\n", d.Body)
			err = d.Ack(true)
			if err != nil {
				fmt.Printf("Mack err: %s", err)
			}
		}
	}()

	time.Sleep(time.Hour)
}

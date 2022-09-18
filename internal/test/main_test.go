package test

import (
	"context"
	"fmt"
	"testing"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/jdxj/cyber-wagon/config"
	gatewayM "github.com/jdxj/cyber-wagon/internal/gateway/model"
	logicM "github.com/jdxj/cyber-wagon/internal/logic/model"
	"github.com/jdxj/cyber-wagon/internal/util"
)

func TestLogicAndPushMQ(t *testing.T) {
	config.Init("../../config/test.yaml")
	util.InitMQ(config.GetRabbitmq())
	logicM.InitCh("test-logic")

	gatewayName := "test-gateway"
	gatewayM.InitQueue(gatewayName)

	go func() {
		for {
			err := gatewayM.PublishLogic(context.Background(), amqp.Publishing{
				Body: []byte("abc"),
			})
			if err != nil {
				fmt.Printf("publish logic err: %s\n", err)
			}
			time.Sleep(time.Second)
		}
	}()

	lmd := logicM.ConsumeLogic()
	go func() {
		for d := range lmd {
			fmt.Printf("consume logic: %s\n", d.Body)
			err := d.Ack(false)
			if err != nil {
				fmt.Printf("ack consume logic err: %s\n", err)
			}
		}
	}()

	go func() {
		for {
			err := logicM.PublishPush(context.Background(), amqp.Publishing{
				Body: []byte("def"),
			})
			if err != nil {
				fmt.Printf("publish push err: %s\n", err)
			}
			time.Sleep(time.Second)
		}
	}()

	gmd := gatewayM.ConsumePush()
	go func() {
		for d := range gmd {
			fmt.Printf("consume push delivery tag: %d\n", d.DeliveryTag)
			fmt.Printf("consume push: %s\n", d.Body)
			err := d.Ack(false)
			if err != nil {
				fmt.Printf("ack consume push err: %s\n", err)
			}
		}

		fmt.Printf("----------------------------------------------")
	}()

	go func() {
		for {
			err := gatewayM.PublishGateway(context.Background(), gatewayName+"_bk", amqp.Publishing{
				Body: []byte("test-self-send"),
			})
			if err != nil {
				fmt.Printf("publish gateway err: %s\n", err)
			}
			time.Sleep(time.Second)
		}
	}()

	gatewayD := gatewayM.ConsumeGateway()
	go func() {
		for d := range gatewayD {
			fmt.Printf("consume gateway delivery tag: %d\n", d.DeliveryTag)
			fmt.Printf("consume gateway: %s\n", d.Body)
			err := d.Ack(true)
			if err != nil {
				fmt.Printf("ack consume gateway err: %s\n", err)
			}
		}
	}()

	time.Sleep(time.Hour)
}

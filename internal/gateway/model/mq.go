package model

import (
	"context"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"

	"github.com/jdxj/cyber-wagon/internal/util"
)

var (
	gatewayName       string
	gatewayQueue      string
	gatewayBindingKey string
)

func InitQueue(name string) {
	ch := util.NewChannel()
	defer ch.Close()

	gatewayName = name
	gatewayQueue = name + "_q"
	gatewayBindingKey = name + "_bk"
	util.CreateQueue(ch, gatewayQueue, gatewayBindingKey)
}

func PublishLogic(ctx context.Context, msg amqp.Publishing) error {
	ch := util.NewChannel()
	return ch.PublishWithContext(ctx, util.Exchange, util.LogicBindingKey,
		false, false, msg)
}

func ConsumePush() <-chan amqp.Delivery {
	ch := util.NewChannel()
	deliveryCh, err := ch.Consume(util.PushQueue, gatewayName,
		false, false, false, false, nil)
	if err != nil {
		logrus.Fatalln(err)
	}
	return deliveryCh
}

func PublishGateway(ctx context.Context, bindingKey string, msg amqp.Publishing) error {
	ch := util.NewChannel()
	return ch.PublishWithContext(ctx, util.Exchange, bindingKey,
		false, false, msg)
}

func ConsumeGateway() <-chan amqp.Delivery {
	ch := util.NewChannel()
	deliveryCh, err := ch.Consume(gatewayQueue, gatewayName,
		false, false, false, false, nil)
	if err != nil {
		logrus.Fatalln(err)
	}
	return deliveryCh
}

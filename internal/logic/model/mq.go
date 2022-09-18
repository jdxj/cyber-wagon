package model

import (
	"context"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"

	"github.com/jdxj/cyber-wagon/internal/util"
)

var (
	logicName string
)

func InitCh(name string) {
	logicName = name
}

func ConsumeLogic() <-chan amqp.Delivery {
	ch := util.NewChannel()
	deliveryCh, err := ch.Consume(util.LogicQueue, logicName,
		false, false, false, false, nil)
	if err != nil {
		logrus.Fatalln(err)
	}
	return deliveryCh
}

func PublishPush(ctx context.Context, msg amqp.Publishing) error {
	ch := util.NewChannel()
	return ch.PublishWithContext(ctx, util.Exchange, util.PushBindingKey,
		false, false, msg)
}

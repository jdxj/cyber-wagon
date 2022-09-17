package util

import (
	"context"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"

	"github.com/jdxj/cyber-wagon/config"
)

const (
	Exchange = "im"
	Kind     = "topic"

	LogicQueue      = "logic_q"
	LogicBindingKey = "logic_bk"
)

var (
	MQ *amqp.Connection
)

func InitMQ(cfg config.Rabbitmq) {
	dsn := fmt.Sprintf("amqp://%s:%s@%s:%d", cfg.User, cfg.Pass, cfg.Host, cfg.Port)
	conn, err := amqp.Dial(dsn)
	if err != nil {
		logrus.Fatalln(err)
	}
	MQ = conn

	createExchange()
	createLogicQueue()
}

func createExchange() {
	ch, err := MQ.Channel()
	if err != nil {
		logrus.Fatalln(err)
	}
	defer ch.Close()

	err = ch.ExchangeDeclare(Exchange, Kind,
		true, false, false, false, nil)
	if err != nil {
		logrus.Fatalln(err)
	}
}

func createLogicQueue() {
	ch, err := MQ.Channel()
	if err != nil {
		logrus.Fatalln(err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(LogicQueue,
		true, false, false, false, nil)
	if err != nil {
		logrus.Fatalln(err)
	}

	err = ch.QueueBind(q.Name, LogicBindingKey, Exchange, false, nil)
	if err != nil {
		logrus.Fatalln(err)
	}
}

func consume(ch *amqp.Channel, queue, consumer string) (<-chan amqp.Delivery, error) {
	return ch.Consume(queue, consumer,
		false, false, false, false, nil)
}

func publish(ctx context.Context, ch *amqp.Channel, routingKey string, msg amqp.Publishing) error {
	return ch.PublishWithContext(ctx, Exchange,
		routingKey, false, false, msg)
}

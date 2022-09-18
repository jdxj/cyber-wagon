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

	PushQueue      = "push_q"
	PushBindingKey = "push_bk"
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

	ch, err := MQ.Channel()
	if err != nil {
		logrus.Fatalln(err)
	}
	defer ch.Close()

	CreateExchange(ch, Exchange, Kind)
	CreateQueue(ch, LogicQueue, LogicBindingKey)
	CreateQueue(ch, PushQueue, PushBindingKey)
}

func NewChannel() *amqp.Channel {
	ch, err := MQ.Channel()
	if err != nil {
		logrus.Fatalln(err)
	}
	return ch
}

func CreateExchange(ch *amqp.Channel, exchange, kind string) {
	err := ch.ExchangeDeclare(exchange, kind,
		true, false, false, false, nil)
	if err != nil {
		logrus.Fatalln(err)
	}
}

func CreateQueue(ch *amqp.Channel, queue, bindingKey string) {
	q, err := ch.QueueDeclare(queue,
		true, false, false, false, nil)
	if err != nil {
		logrus.Fatalln(err)
	}

	err = ch.QueueBind(q.Name, bindingKey, Exchange, false, nil)
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

package imageProcessing

import (
	"github.com/pkg/errors"
	"github.com/streadway/amqp"
	"paint/pkg/rabbitmq"
)

type AMQPConfig struct {
	Create struct {
		ExchangeName string
		ExchangeType string
		RoutingKey   string
		QueueName    string
	}
}

type AMQP struct {
	config   AMQPConfig
	rabbitmq *rabbitmq.RabbitMQ
}

func NewAMQP(config AMQPConfig, rabbitmq *rabbitmq.RabbitMQ) AMQP {
	return AMQP{
		config:   config,
		rabbitmq: rabbitmq,
	}
}

func (a AMQP) Setup() error {
	channel, err := a.rabbitmq.Channel()
	if err != nil {
		return errors.Wrap(err, "failed to open channel")
	}
	defer channel.Close()

	if err := a.declareCreate(channel); err != nil {
		return err
	}

	return nil
}

func (a AMQP) declareCreate(channel *amqp.Channel) error {
	if err := channel.ExchangeDeclare(
		a.config.Create.ExchangeName,
		a.config.Create.ExchangeType,
		true,
		false,
		false,
		false,
		nil,
	); err != nil {
		return errors.Wrap(err, "failed to declare exchange")
	}

	if _, err := channel.QueueDeclare(
		a.config.Create.QueueName,
		true,
		false,
		false,
		false,
		amqp.Table{"x-queue-mode": "lazy"},
	); err != nil {
		return errors.Wrap(err, "failed to declare queue")
	}

	if err := channel.QueueBind(
		a.config.Create.QueueName,
		a.config.Create.RoutingKey,
		a.config.Create.ExchangeName,
		false,
		nil,
	); err != nil {
		return errors.Wrap(err, "failed to bind queue")
	}

	return nil
}

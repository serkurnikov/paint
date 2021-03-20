package rabbitMQ

import (
	"context"
	"paint/pkg/rabbitmq"
)

type Ctx = context.Context

type RabbitMQ struct {
	*rabbitmq.RabbitMQ
}

func New(ctx Ctx) (_ *RabbitMQ, err error) {
	return nil, err
}


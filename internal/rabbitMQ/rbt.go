package rabbitMQ

import (
	"context"
	"log"
	"paint/pkg/rabbitmq"
	"time"
)

type Ctx = context.Context

type RabbitMQ struct {
	*rabbitmq.RabbitMQ
}

func New(ctx Ctx) (_ *RabbitMQ, err error) {
	/*returnErrs := []error{
		app.ErrRabbitConnect,
	}*/

	r := &RabbitMQ{}
	config := rabbitmq.Config{
		Schema:               "",
		Username:             "",
		Password:             "",
		Host:                 "",
		Port:                 0,
		Vhost:                "",
		ConnectionName:       "",
		ChannelNotifyTimeout: 0,
		Reconnect: struct {
			Interval   time.Duration
			MaxAttempt int
		}{},
	}

	r.RabbitMQ, err = rabbitmq.New(config)
	if err := r.Connect(); err != nil {log.Fatalln(err)}


	/*
	processingAMQP := imageProcessing.NewAMQP(s.cfg.ImageProcessingAMQP, rbt)
	if err := processingAMQP.Setup(); err != nil {log.Fatalln(err)}

	rabbitImageProcessing := imageProcessing.NewCreate(rbt)
	*/


	if err != nil {
		return nil, err
	}

	return r, err
}


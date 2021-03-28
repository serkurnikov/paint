package rabbitmq
//http://www.inanzzz.com/index.php/post/0aeg/creating-a-rabbitmq-producer-example-with-golang

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/streadway/amqp"
)

type Config struct {
	Schema               string
	Username             string
	Password             string
	Host                 string
	Port                 int
	Vhost                string
	ConnectionName       string
	ChannelNotifyTimeout time.Duration
	Reconnect            struct {
		Interval   time.Duration
		MaxAttempt int
	}
}

type RabbitMQ struct {
	mux                  sync.RWMutex
	config               Config
	dialConfig           amqp.Config
	connection           *amqp.Connection
	ChannelNotifyTimeout time.Duration
}

func New(config Config) (*RabbitMQ, error) {
	return &RabbitMQ{
		config:               config,
		dialConfig:           amqp.Config{Properties: amqp.Table{"connection_name": config.ConnectionName}},
		ChannelNotifyTimeout: config.ChannelNotifyTimeout,
	}, nil
}

// Connect creates a new connection. Use once at application
// startup.
func (r *RabbitMQ) Connect() error {
	con, err := amqp.DialConfig(fmt.Sprintf(
		"%s://%s:%s@%s:%v/%s",
		r.config.Schema,
		r.config.Username,
		r.config.Password,
		r.config.Host,
		r.config.Port,
		r.config.Vhost,
	), r.dialConfig)
	if err != nil {
		return err
	}

	r.connection = con

	go r.reconnect()

	return nil
}

// Channel returns a new `*amqp.Channel` instance. You must
// call `defer channel.Close()` as soon as you obtain one.
// Sometimes the connection might be closed unintentionally so
// as a graceful handling, try to connect only once.
func (r *RabbitMQ) Channel() (*amqp.Channel, error) {
	if r.connection == nil {
		if err := r.Connect(); err != nil {
			return nil, errors.New("connection is not open")
		}
	}

	channel, err := r.connection.Channel()
	if err != nil {
		return nil, err
	}

	return channel, nil
}

// Connection exposes the essentials of the current connection.
// You should not normally use this but it is there for special
// use cases.
func (r *RabbitMQ) Connection() *amqp.Connection {
	return r.connection
}

// Shutdown triggers a normal shutdown. Use this when you wish
// to shutdown your current connection or if you are shutting
// down the application.
func (r *RabbitMQ) Shutdown() error {
	if r.connection != nil {
		return r.connection.Close()
	}

	return nil
}

// reconnect reconnects to server if the connection or a channel
// is closed unexpectedly. Normal shutdown is ignored. It tries
// maximum of 7200 times and sleeps half a second in between
// each try which equals to 1 hour.
func (r *RabbitMQ) reconnect() {
WATCH:

	conErr := <-r.connection.NotifyClose(make(chan *amqp.Error))
	if conErr != nil {
		log.Println("CRITICAL: Connection dropped, reconnecting")

		var err error

		for i := 1; i <= r.config.Reconnect.MaxAttempt; i++ {
			r.mux.RLock()
			r.connection, err = amqp.DialConfig(fmt.Sprintf(
				"%s://%s:%s@%s:%v/%s",
				r.config.Schema,
				r.config.Username,
				r.config.Password,
				r.config.Host,
				r.config.Port,
				r.config.Vhost,
			), r.dialConfig)
			r.mux.RUnlock()

			if err == nil {
				log.Println("INFO: Reconnected")

				goto WATCH
			}

			time.Sleep(r.config.Reconnect.Interval)
		}

		log.Println(errors.Wrap(err, "CRITICAL: Failed to reconnect"))
	} else {
		log.Println("INFO: Connection dropped normally, will not reconnect")
	}
}

// Package config provides configurations for subcommands.
//
// Default values can be obtained from various sources (constants,
// environment variables, etc.) and then overridden by flags.
//
// As configuration is global you can get it only once for safety:
// you can call only one of Get… functions and call it just once.
package config

import (
	"github.com/go-sql-driver/mysql"
	"github.com/powerman/appcfg"
	"github.com/spf13/pflag"
	"paint/internal/app/imageProcessing"
	"paint/pkg/def"
	"paint/pkg/netx"
	"paint/pkg/rabbitmq"
	"time"
)

// EnvPrefix defines common prefix for environment variables.
const envPrefix = "EXAMPLE_"

// All configurable values of the microservice.
//
// If microservice may runs in different ways (e.g. using CLI subcommands)
// then these subcommands may use subset of these values.
var all = &struct { //nolint:gochecknoglobals // Config is global anyway.
	APIKeyAdmin                   appcfg.NotEmptyString `env:"APIKEY_ADMIN"`
	AddrHost                      appcfg.NotEmptyString `env:"ADDR_HOST"`
	AddrPort                      appcfg.Port           `env:"ADDR_PORT"`
	MetricsAddrPort               appcfg.Port           `env:"METRICS_ADDR_PORT"`
	MySQLAddrHost                 appcfg.NotEmptyString `env:"MYSQL_ADDR_HOST"`
	MySQLAddrPort                 appcfg.Port           `env:"MYSQL_ADDR_PORT"`
	MySQLAuthLogin                appcfg.NotEmptyString `env:"MYSQL_AUTH_LOGIN"`
	MySQLAuthPass                 appcfg.String         `env:"MYSQL_AUTH_PASS"`
	MySQLDBName                   appcfg.NotEmptyString `env:"MYSQL_DB"`
	MySQLGooseDir                 appcfg.NotEmptyString
	RabbitMQSchema                appcfg.NotEmptyString `env:"RABBITMQ_SCHEMA"`
	RabbitMQUserName              appcfg.NotEmptyString `env:"RABBITMQ_USERNAME"`
	RabbitMQPass                  appcfg.NotEmptyString `env:"RABBITMQ_PASS"`
	RabbitMQHost                  appcfg.NotEmptyString `env:"RABBITMQ_HOST"`
	RabbitMQPort                  appcfg.Port           `env:"RABBITMQ_PORT"`
	RabbitMQVHost                 appcfg.NotEmptyString `env:"RABBITMQ_VHOST"`
	RabbitMQVConnectionName       appcfg.NotEmptyString `env:"RABBITMQ_CONNECTION_NAME"`
	RabbitMQVChannelNotifyTimeOut appcfg.Duration       `env:"RABBITMQ_CHANNEL_NOTIFY_TIMEOUT"`
	RabbitMQVReconnectInterval    appcfg.Duration       `env:"RABBITMQ_RECONNECT_INTERVAL"`
	RabbitMQVReconnectMaxAttempt  appcfg.Int            `env:"RABBITMQ_RECONNECT_MAX_ATTEMPT"`
	RabbitMQExchangeName          appcfg.NotEmptyString `env:"RABBITMQ_EXCHANGE_NAME"`
	RabbitMQExchangeType          appcfg.NotEmptyString `env:"RABBITMQ_EXCHANGE_TYPE"`
	RabbitMQRoutingKey            appcfg.NotEmptyString `env:"RABBITMQ_ROUTING_KEY"`
	RabbitMQQueueName             appcfg.NotEmptyString `env:"RABBITMQ_QUEUE_NAME"`
}{ // Defaults, if any:
	AddrHost:        appcfg.MustNotEmptyString(def.Hostname),
	AddrPort:        appcfg.MustPort("8000"),
	MetricsAddrPort: appcfg.MustPort("9000"),
	MySQLAddrPort:   appcfg.MustPort("3306"),
	MySQLAuthLogin:  appcfg.MustNotEmptyString(def.ProgName),
	MySQLDBName:     appcfg.MustNotEmptyString(def.ProgName),
	MySQLGooseDir:   appcfg.MustNotEmptyString("internal/migrations/mysql"),
}

// FlagSets for all CLI subcommands which use flags to set config values.
type FlagSets struct {
	Serve      *pflag.FlagSet
	GooseMySQL *pflag.FlagSet
	RabbitMQ   *pflag.FlagSet
}

var fs FlagSets //nolint:gochecknoglobals // Flags are global anyway.

// Init updates config defaults (from env) and setup subcommands flags.
//
// Init must be called once before using this package.
func Init(flagsets FlagSets) error {
	fs = flagsets

	fromEnv := appcfg.NewFromEnv(envPrefix)
	err := appcfg.ProvideStruct(all, fromEnv)
	if err != nil {
		return err
	}

	appcfg.AddPFlag(fs.Serve, &all.AddrHost, "host", "host to serve OpenAPI")
	appcfg.AddPFlag(fs.Serve, &all.AddrPort, "port", "port to serve OpenAPI")
	appcfg.AddPFlag(fs.Serve, &all.MetricsAddrPort, "metrics.port", "port to serve Prometheus metrics")
	appcfg.AddPFlag(fs.Serve, &all.MySQLAddrHost, "mysql.host", "host to connect to MySQL")
	appcfg.AddPFlag(fs.Serve, &all.MySQLAddrPort, "mysql.port", "port to connect to MySQL")
	appcfg.AddPFlag(fs.Serve, &all.MySQLAuthLogin, "mysql.user", "MySQL username")
	appcfg.AddPFlag(fs.Serve, &all.MySQLAuthPass, "mysql.pass", "MySQL password")
	appcfg.AddPFlag(fs.Serve, &all.MySQLDBName, "mysql.dbname", "MySQL database name")

	appcfg.AddPFlag(fs.GooseMySQL, &all.MySQLAddrHost, "mysql.host", "host to connect to MySQL")
	appcfg.AddPFlag(fs.GooseMySQL, &all.MySQLAddrPort, "mysql.port", "port to connect to MySQL")
	appcfg.AddPFlag(fs.GooseMySQL, &all.MySQLAuthLogin, "mysql.user", "MySQL username")
	appcfg.AddPFlag(fs.GooseMySQL, &all.MySQLAuthPass, "mysql.pass", "MySQL password")
	appcfg.AddPFlag(fs.GooseMySQL, &all.MySQLDBName, "mysql.dbname", "MySQL database name")

	appcfg.AddPFlag(fs.RabbitMQ, &all.RabbitMQSchema, "rabbitmq.schema", "Rabbit schema")
	appcfg.AddPFlag(fs.RabbitMQ, &all.RabbitMQUserName, "rabbitmq.userName", "Rabbit username")
	appcfg.AddPFlag(fs.RabbitMQ, &all.RabbitMQPass, "rabbitmq.pass", "Rabbit password")
	appcfg.AddPFlag(fs.RabbitMQ, &all.RabbitMQHost, "rabbitmq.host", "host to connect to Rabbit")
	appcfg.AddPFlag(fs.RabbitMQ, &all.RabbitMQPort, "rabbitmq.port", "port to connect to Rabbit")
	appcfg.AddPFlag(fs.RabbitMQ, &all.RabbitMQVHost, "rabbitmq.VHost", "vhost to connect to Rabbit")
	appcfg.AddPFlag(fs.RabbitMQ, &all.RabbitMQVConnectionName, "rabbitmq.connectionName", "Rabbit")
	appcfg.AddPFlag(fs.RabbitMQ, &all.RabbitMQVChannelNotifyTimeOut, "rabbitmq.channelNotifyTimeOut", "Rabbit")
	appcfg.AddPFlag(fs.RabbitMQ, &all.RabbitMQVReconnectInterval, "rabbitmq.reconnectInterval", "Rabbit")
	appcfg.AddPFlag(fs.RabbitMQ, &all.RabbitMQVReconnectMaxAttempt, "rabbitmq.reconnectMaxAttempt", "Rabbit")

	appcfg.AddPFlag(fs.RabbitMQ, &all.RabbitMQExchangeName, "rabbitmq.exchangeName", "Rabbit exchangeName")
	appcfg.AddPFlag(fs.RabbitMQ, &all.RabbitMQExchangeType, "rabbitmq.exchangeType", "Rabbit exchangeType")
	appcfg.AddPFlag(fs.RabbitMQ, &all.RabbitMQRoutingKey, "rabbitmq.routingKey", "Rabbit routingKey")
	appcfg.AddPFlag(fs.RabbitMQ, &all.RabbitMQQueueName, "rabbitmq.queueName", "Rabbit queueName")

	return nil
}

// ServeConfig contains configuration for subcommand.
type ServeConfig struct {
	APIKeyAdmin         string
	Addr                netx.Addr
	MetricsAddr         netx.Addr
	MySQL               *mysql.Config
	MySQLGooseDir       string
	RabbitMQ            rabbitmq.Config
	ImageProcessingAMQP imageProcessing.AMQPConfig
}

// GetServe validates and returns configuration for subcommand.
func GetServe() (c *ServeConfig, err error) {
	defer cleanup()

	c = &ServeConfig{
		APIKeyAdmin:   all.APIKeyAdmin.Value(&err),
		Addr:          netx.NewAddr(all.AddrHost.Value(&err), all.AddrPort.Value(&err)),
		MetricsAddr:   netx.NewAddr(all.AddrHost.Value(&err), all.MetricsAddrPort.Value(&err)),
		MySQLGooseDir: all.MySQLGooseDir.Value(&err),
		RabbitMQ: rabbitmq.Config{
			Schema:               all.RabbitMQSchema.Value(&err),
			Username:             all.RabbitMQUserName.Value(&err),
			Password:             all.RabbitMQPass.Value(&err),
			Host:                 all.RabbitMQHost.Value(&err),
			Port:                 all.RabbitMQPort.Value(&err),
			Vhost:                all.RabbitMQVHost.Value(&err),
			ConnectionName:       all.RabbitMQVConnectionName.Value(&err),
			ChannelNotifyTimeout: all.RabbitMQVChannelNotifyTimeOut.Value(&err),
			Reconnect: struct {
				Interval   time.Duration
				MaxAttempt int
			}{all.RabbitMQVReconnectInterval.Value(&err), all.RabbitMQVReconnectMaxAttempt.Value(&err)},
		},
		ImageProcessingAMQP: imageProcessing.AMQPConfig{Create: struct {
			ExchangeName string
			ExchangeType string
			RoutingKey   string
			QueueName    string
		}{ExchangeName: all.RabbitMQExchangeName.Value(&err),
			ExchangeType: all.RabbitMQExchangeType.Value(&err),
			RoutingKey:   all.RabbitMQRoutingKey.Value(&err),
			QueueName:    all.RabbitMQQueueName.Value(&err)}},
	}
	if err != nil {
		return nil, appcfg.WrapPErr(err, fs.Serve, all)
	}
	return c, nil
}

// Cleanup must be called by all Get* functions to ensure second call to
// any of them will panic.
func cleanup() {
	all = nil
}

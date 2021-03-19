package config

import (
	"github.com/powerman/appcfg"
	"github.com/spf13/pflag"
	"paint/internal/app/imageProcessing"
	"paint/pkg/cobrax"
	"paint/pkg/def"
	"paint/pkg/netx"
	"paint/pkg/rabbitmq"
	"time"
)

const envPrefix = "EXAMPLE_"

var all = &struct {
	APIKeyAdmin                   appcfg.NotEmptyString `env:"APIKEY_ADMIN"`
	AddrHost                      appcfg.NotEmptyString `env:"ADDR_HOST"`
	AddrPort                      appcfg.Port           `env:"ADDR_PORT"`
	MetricsAddrPort               appcfg.Port           `env:"METRICS_ADDR_PORT"`
	SqlAddrHost                   appcfg.NotEmptyString `env:"SQL_ADDR_HOST"`
	SqlAddrPort                   appcfg.Port           `env:"SQL_ADDR_PORT"`
	SqlAuthLogin                  appcfg.NotEmptyString `env:"SQL_AUTH_LOGIN"`
	SqlAuthPass                   appcfg.String         `env:"SQL_AUTH_PASS"`
	SqlDbName                     appcfg.NotEmptyString `env:"SQL_DB"`
	SqlGooseDir                   appcfg.NotEmptyString
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
}{
	AddrHost:        appcfg.MustNotEmptyString(def.Hostname),
	AddrPort:        appcfg.MustPort("8000"),
	MetricsAddrPort: appcfg.MustPort("9000"),
	SqlAddrPort:     appcfg.MustPort("3306"),
	SqlAuthLogin:    appcfg.MustNotEmptyString(def.ProgName),
	SqlDbName:       appcfg.MustNotEmptyString(def.ProgName),
	SqlGooseDir:     appcfg.MustNotEmptyString("internal/migrations/mysql"),
}

type FlagSets struct {
	Serve    *pflag.FlagSet
	GooseSQL *pflag.FlagSet
	RabbitMQ *pflag.FlagSet
}

var fs FlagSets

func Init(flagSets FlagSets) error {
	fs = flagSets

	fromEnv := appcfg.NewFromEnv(envPrefix)
	err := appcfg.ProvideStruct(all, fromEnv)
	if err != nil {
		return err
	}

	appcfg.AddPFlag(fs.Serve, &all.AddrHost, "host", "host to serve OpenAPI")
	appcfg.AddPFlag(fs.Serve, &all.AddrPort, "port", "port to serve OpenAPI")
	appcfg.AddPFlag(fs.Serve, &all.MetricsAddrPort, "metrics.port", "port to serve Prometheus metrics")
	appcfg.AddPFlag(fs.Serve, &all.SqlAddrHost, "sql.host", "host to connect to SQL")
	appcfg.AddPFlag(fs.Serve, &all.SqlAddrPort, "sql.port", "port to connect to SQL")
	appcfg.AddPFlag(fs.Serve, &all.SqlAuthLogin, "sql.user", "SQL username")
	appcfg.AddPFlag(fs.Serve, &all.SqlAuthPass, "sql.pass", "SQL password")
	appcfg.AddPFlag(fs.Serve, &all.SqlDbName, "sql.dbname", "SQL database name")

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

type ServeConfig struct {
	APIKeyAdmin         string
	Addr                netx.Addr
	MetricsAddr         netx.Addr
	SQLGooseDir         string
	RabbitMQ            rabbitmq.Config
	ImageProcessingAMQP imageProcessing.AMQPConfig
}

func GetServe() (c *ServeConfig, err error) {
	defer cleanup()

	c = &ServeConfig{
		APIKeyAdmin: all.APIKeyAdmin.Value(&err),
		Addr:        netx.NewAddr(all.AddrHost.Value(&err), all.AddrPort.Value(&err)),
		MetricsAddr: netx.NewAddr(all.AddrHost.Value(&err), all.MetricsAddrPort.Value(&err)),
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

func GetGooseSQL() (c *cobrax.GooseSQLConfig, err error) {
	defer cleanup()

	c = &cobrax.GooseSQLConfig{
		SQL: def.NewSQLConfig(def.SQLConfig{
			Addr: netx.NewAddr(all.SqlAddrHost.Value(&err), all.SqlAddrPort.Value(&err)),
			User: all.SqlAuthLogin.Value(&err),
			Pass: all.SqlAuthPass.Value(&err),
			DB:   all.SqlDbName.Value(&err),
		}),
		SQLGooseDir: all.SqlGooseDir.Value(&err),
	}
	if err != nil {
		return nil, appcfg.WrapPErr(err, fs.GooseSQL, all)
	}
	return c, nil
}

func cleanup() {
	all = nil
}

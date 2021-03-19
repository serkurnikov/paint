package def

import (
	"github.com/go-sql-driver/mysql"
	"github.com/powerman/go-service-example/pkg/netx"
)

// SQLConfig contains SQL connection and authentication details.
type SQLConfig struct {
	Addr netx.Addr
	User string
	Pass string
	DB   string
}

// NewSQLConfig creates a new default config for SQL.
func NewSQLConfig(cfg SQLConfig) *mysql.Config {
	c := mysql.NewConfig()
	c.User = cfg.User
	c.Passwd = cfg.Pass
	c.Net = "tcp"
	c.Addr = cfg.Addr.String()
	c.DBName = cfg.DB
	c.Params = map[string]string{
		"sql_mode": "'TRADITIONAL'", // 5.6 defaults + all strict modes.
	}
	c.Collation = "utf8mb4_unicode_ci"
	c.ParseTime = true
	c.RejectReadOnly = true
	return c
}

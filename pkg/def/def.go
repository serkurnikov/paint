package def

import (
	"time"

	"github.com/powerman/must"
	"github.com/powerman/structlog"
)

// Log field names.
const (
	LogRemote     = "remote"
	LogFunc       = "func"
	LogGRPCCode   = "grpcCode"
	LogServer     = "server"
	LogAddr       = "addr"       // host:port.
	LogHost       = "host"       // DNS hostname or IPv4/IPv6 address.
	LogPort       = "port"       // TCP/UDP port number.
	LogHTTPMethod = "httpMethod" // GET, POST, etc.
	LogHTTPStatus = "httpStatus" // Status code: 200, 404, etc.

)

// Init must be called once before using this package.
// It provides common initialization for both commands and tests.
func Init() {
	must.AbortIf = must.PanicIf

	structlog.DefaultLogger.
		SetPrefixKeys(
			LogFunc, structlog.KeyUnit,
		).
		SetDefaultKeyvals(
			structlog.KeyTime, structlog.Auto,
		).
		SetSuffixKeys(
			structlog.KeyStack, structlog.KeySource,
		).
		SetKeysFormat(map[string]string{
			structlog.KeyTime:   " %[2]s",
			structlog.KeyStack:  " %6[2]s",
			structlog.KeySource: " %6[2]s",
			structlog.KeyUnit:   " %6[2]s",
			LogGRPCCode:         " %-16.16[2]s",
			LogFunc:             " %[2]s:",
			"duration":          " %[2]q",
			"request":           " %[1]s=% [2]X",
			"response":          " %[1]s=% [2]X",
		}).SetTimeFormat(time.StampMicro)
}

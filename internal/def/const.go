package def

import (
	"runtime/debug"
)

var (
	ver string
)

func Version() string {
	if bi, ok := debug.ReadBuildInfo(); ok {
		if bi.Main.Version == "(devel)" && ver != "" {
			return ver
		}
		return bi.Main.Version
	}
	return "(test)"
}

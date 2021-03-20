// Package migrations provides goose migrations.
package migrations

import (
	goosePkg "github.com/powerman/goose/v2"
	"paint/pkg/def"
)

var goose = def.NewGoose()

func Goose() *goosePkg.Instance { return goose }

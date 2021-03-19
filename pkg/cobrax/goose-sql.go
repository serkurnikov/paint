package cobrax

import (
	"context"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/powerman/go-service-example/pkg/migrate"
	goosePkg "github.com/powerman/goose/v2"
	"github.com/spf13/cobra"
	"strings"
)

// GooseSQLConfig contain configuration for goose command.
type GooseSQLConfig struct {
	SQL         *mysql.Config
	SQLGooseDir string
}

// NewGooseSQLCmd creates new goose command executed by run.
func NewGooseSQLCmd(goose *goosePkg.Instance, getCfg func() (*GooseSQLConfig, error)) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "goose-sql",
		Short: "Migrate SQL database schema",
		Args:  gooseArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			gooseCmd := strings.Join(args, " ")

			ctx := context.Background()
			cfg, err := getCfg()
			if err != nil {
				return fmt.Errorf("failed to get config: %w", err)
			}

			err = migrate.Run(ctx, goose, cfg.SQLGooseDir, gooseCmd, cfg.SQL)
			if err != nil {
				return fmt.Errorf("failed to run goose %s: %w", gooseCmd, err)
			}
			return nil
		},
	}
	cmd.SetUsageTemplate(gooseUsageTemplate)
	return cmd
}

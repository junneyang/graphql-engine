package commands

import (
	"time"

	"github.com/hasura/graphql-engine/cli"
	mig "github.com/hasura/graphql-engine/cli/migrate/cmd"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func newMigrateCreateCmd(ec *cli.ExecutionContext) *cobra.Command {
	opts := &migrateCreateOptions{
		EC: ec,
	}

	migrateCreateCmd := &cobra.Command{
		Use:          "create [migration-name]",
		Short:        "Create files required for a migration",
		Long:         "Create sql and yaml files required for a migration",
		SilenceUsage: true,
		Args:         cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.name = args[0]
			return opts.run()
		},
	}

	return migrateCreateCmd
}

type migrateCreateOptions struct {
	EC *cli.ExecutionContext

	name string
}

func (o *migrateCreateOptions) run() error {
	timestamp := getTime()
	err := mig.CreateCmd(o.EC.MigrationDir, timestamp, o.name)
	if err != nil {
		return errors.Wrap(err, "error creating migration files")
	}
	o.EC.Logger.Infof("Migration files created with version %d_%s.[up|down].[yaml|sql]", timestamp, o.name)
	return nil
}

func getTime() int64 {
	startTime := time.Now()
	return startTime.UnixNano() / int64(time.Millisecond)
}

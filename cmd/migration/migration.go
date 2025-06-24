package migration

import "github.com/spf13/cobra"

func Command() *cobra.Command {
	commands := &cobra.Command{
		Use:   "migration",
		Short: "for creating, run, or rollback migrations",
	}

	commands.AddCommand(
		createCommand(),
		migrateCommand(),
		rollbackCommand(),
	)

	return commands
}

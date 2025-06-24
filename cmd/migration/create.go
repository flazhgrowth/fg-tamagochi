package migration

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

func createCommand() *cobra.Command {
	commands := &cobra.Command{
		Use:   "create",
		Run:   create,
		Short: "create new migration",
	}
	commands.Flags().String("path", "", "path to save created migration file")
	commands.Flags().String("name", "", "migration name")

	return commands
}

func create(cmd *cobra.Command, args []string) {
	// ex: migrate create -ext sql -dir ./path/to/save/the/migration -seq migration_name
	path, err := cmd.Flags().GetString("path")
	if err != nil {
		panic(err)
	}
	migrationName, err := cmd.Flags().GetString("name")
	if err != nil {
		panic(err)
	}
	execCmd := exec.Command("migrate", "create", "-ext", "sql", "-dir", path, "-seq", migrationName)
	res, err := execCmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s created\n%s\n", migrationName, string(res))
}

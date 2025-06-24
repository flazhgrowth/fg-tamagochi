package migration

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

func migrateCommand() *cobra.Command {
	commands := &cobra.Command{
		Use:   "migrate",
		Run:   migrate,
		Short: "migrate all the migrations",
	}
	commands.Flags().String("path", "", "migration path")
	commands.Flags().String("database", "", "database dsn")

	return commands
}

func migrate(cmd *cobra.Command, args []string) {
	// migrate -path ./path/to/migration/files -database dsn up
	path, err := cmd.Flags().GetString("path")
	if err != nil {
		panic(err)
	}
	dsn, err := cmd.Flags().GetString("database")
	if err != nil {
		panic(err)
	}
	fmt.Println("about to run: ", "migrate", "-path", path, "-database", fmt.Sprintf("'%s'", dsn), "up")
	execCmd := exec.Command("migrate", "-path", path, "-database", dsn, "up")
	res, err := execCmd.CombinedOutput()
	if err != nil {
		panic(fmt.Sprintf("error on executing migrate up: %s", err.Error()))
	}
	fmt.Println("migration migrated:", string(res))
}

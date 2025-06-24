package migration

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

func rollbackCommand() *cobra.Command {
	commands := &cobra.Command{
		Use:   "rollback",
		Run:   rollback,
		Short: "rollback(migrate down) all the migrations",
	}
	commands.Flags().String("path", "", "migration path")
	commands.Flags().String("database", "", "database dsn")

	return commands
}

func rollback(cmd *cobra.Command, args []string) {
	// migrate -path ./path/to/migration/files -database dsn down -all
	path, err := cmd.Flags().GetString("path")
	if err != nil {
		panic(err)
	}
	dsn, err := cmd.Flags().GetString("database")
	if err != nil {
		panic(err)
	}
	fmt.Println("about to run: ", "migrate", "-path", path, "-database", fmt.Sprintf("'%s'", dsn), "down", "-all")
	execCmd := exec.Command("migrate", "-path", path, "-database", dsn, "down", "-all")
	res, err := execCmd.CombinedOutput()
	if err != nil {
		panic(fmt.Sprintf("error on executing migrate down: %s", err.Error()))
	}
	fmt.Println("migration rollback:", string(res))
}

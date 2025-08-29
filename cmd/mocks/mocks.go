package mocks

import (
	"os/exec"

	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	commands := &cobra.Command{
		Use:   "mocks",
		Short: "initialize app mocks",
		Run: func(cmd *cobra.Command, args []string) {
			generateMocks()
		},
	}

	return commands
}

func generateMocks() {
	cmd := exec.Command("mockery")
	resp, err := cmd.CombinedOutput()
	if err != nil {
		if resp != nil {
			println(string(resp))
		}
		panic(err)
	}
	println(string(resp))
	println("Done generating mocks...")
}

package docs

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	commands := &cobra.Command{
		Use:   "docs",
		Short: "initialize app swagger docs",
		Run: func(cmd *cobra.Command, args []string) {
			generateDocs()
		},
	}

	return commands
}

func generateDocs() {
	// swag init -g main.go --dir ./exampleapp --output ./exampleapp/docs
	fmt.Println("Generating docs...")
	cmd := exec.Command("swag", "init", "-g", "main.go", "--dir", "./", "--output", "./docs")
	resp, err := cmd.CombinedOutput()
	if err != nil {
		if resp != nil {
			fmt.Println(string(resp))
		}
		panic(err)
	}
	fmt.Println(string(resp))
	fmt.Println("Done generating docs...")
}

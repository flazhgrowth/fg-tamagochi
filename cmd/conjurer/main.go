package main

import (
	"github.com/flazhgrowth/fg-tamagochi/cmd/docs"
	"github.com/flazhgrowth/fg-tamagochi/cmd/initproject"
	"github.com/flazhgrowth/fg-tamagochi/cmd/migration"
	"github.com/spf13/cobra"
)

func main() {
	conjure()
}

func conjure() {
	root := &cobra.Command{
		Use: "app",
	}

	smith := &cobra.Command{
		Use:   "conjure",
		Short: "conjure a command",
		Long:  "conjure a command",
	}
	smith.AddCommand(
		migration.Command(),
		initproject.Command(),
		docs.Command(),
	)

	root.AddCommand(
		smith,
	)

	if err := root.Execute(); err != nil {
		panic(err)
	}
}

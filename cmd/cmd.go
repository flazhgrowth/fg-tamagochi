package cmd

import (
	"github.com/flazhgrowth/fg-tamagochi/cmd/docs"
	"github.com/flazhgrowth/fg-tamagochi/cmd/initproject"
	"github.com/flazhgrowth/fg-tamagochi/cmd/migration"
	"github.com/flazhgrowth/fg-tamagochi/cmd/serve"
	"github.com/spf13/cobra"
)

func Conjure(args CmdArgs) {
	root := &cobra.Command{
		Use: "app",
	}

	smith := &cobra.Command{
		Use:   "conjure",
		Short: "conjure a command",
		Long:  "conjure a command",
	}
	smith.AddCommand(
		serve.Command(args.ServeCmdArgs),
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

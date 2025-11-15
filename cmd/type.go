package cmd

import (
	"github.com/flazhgrowth/fg-tamagochi/cmd/serve"
	"github.com/spf13/cobra"
)

type CmdArgs struct {
	ServeCmdArgs serve.ServeCmdArgs
	Commands     []*cobra.Command
}

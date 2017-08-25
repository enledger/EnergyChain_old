package main

import (
	"os"
	"github.com/spf13/cobra"
	"github.com/tendermint/tmlibs/cli"
	"github.com/tendermint/cosmos-sdk/cmd/basecoin/commands"
        "github.com/enledger/energychain/energycli"
	"github.com/tendermint/cosmos-sdk/types"
)

func main() {
	var RootCmd = &cobra.Command{
		Use:   "energychain",
		Short: "EnergyChain Blockchain Node",
	}
	RootCmd.AddCommand(
		commands.InitCmd,
		commands.StartCmd,
		commands.UnsafeResetAllCmd,
		commands.VersionCmd,
	)

	commands.RegisterStartPlugin("energychain", func() types.Plugin { return counter.New() })
	cmd := cli.PrepareMainCmd(RootCmd, "CT", os.ExpandEnv("$HOME/.energychain"))
	cmd.Execute()
}

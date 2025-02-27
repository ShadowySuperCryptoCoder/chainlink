package keeper

import (
	"github.com/spf13/cobra"
)

// RootCmd represents the root keeper sub-command to manage keepers
var RootCmd = &cobra.Command{
	Use:   "keeper",
	Short: "Manage keepers",
	Long:  `This command represents a CLI interface to manage keepers.`,
}

func init() {
	RootCmd.AddCommand(deployCmd)
}

package cmd

import (
	"github.com/spf13/cobra"
)

// closeCmd represents the close command
var closeCmd = &cobra.Command{
	Use:   "close",
	Short: "Close the issues of the github repository",
	Long: `Close the issues of the github repository,cli is a CLI tool for Github repositories and issues management,
which will be improving continiously as the awesome tool!`,
	Run: func(cmd *cobra.Command, args []string) {
        Usage()
	},
}

func init() {
	rootCmd.AddCommand(closeCmd)
}

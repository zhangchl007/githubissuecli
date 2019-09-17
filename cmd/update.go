package cmd

import (
	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "update the issue for user repository",
	Long: `update the issue for user repository,gcli is a CLI tool for Github repositories and issues management,
which will be improving continiously as the awesome tool!`,
	Run: func(cmd *cobra.Command, args []string) {
        Usage()
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}

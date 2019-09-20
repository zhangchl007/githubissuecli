package cmd

import (
	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update the repo and issue for the user",
	Long: `Update the repo and issue for the user, gcli is a CLI tool for Github repositories and issues management,
which will be improving continiously as the awesome tool!`,
	Run: func(cmd *cobra.Command, args []string) {
        Usage()
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}

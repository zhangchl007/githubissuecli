package cmd

import (
	"github.com/spf13/cobra"
)

// lockCmd represents the lock command
var lockCmd = &cobra.Command{
	Use:   "lock",
	Short: "Lock the specify issue of repo",
	Long: `Lock the specify issue of repo, gcli is a CLI tool for Github repositories and issues management,
which will be improving continiously as the awesome tool!`,
	Run: func(cmd *cobra.Command, args []string) {
        Usage()
	},
}

func init() {
	rootCmd.AddCommand(lockCmd)

}

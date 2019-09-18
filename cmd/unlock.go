package cmd

import (
	"github.com/spf13/cobra"
)

// unlockCmd represents the unlock command
var unlockCmd = &cobra.Command{
	Use:   "unlock",
	Short: "Unlock the specify issue of repo",
	Long: `Unlock the specify issue of repo, gcli is a CLI tool for Github repositories and issues management,
which will be improving continiously as the awesome tool!`,
	Run: func(cmd *cobra.Command, args []string) {
		Usage()
	},
}

func init() {
	rootCmd.AddCommand(unlockCmd)
}

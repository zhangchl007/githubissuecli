package cmd

import (
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var Title, Body, State, Owner, Label, Name, Description, Homepage string
var Locked,Private bool

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create the repo/issue and the yaml file of issue template",
	Long: `Create the repo/issue and the yaml file of issue template, gcli is a CLI tool for Github repositories and issues management,
which will be improving continiously as the awesome tool!`,
    Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
        Usage()
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

}

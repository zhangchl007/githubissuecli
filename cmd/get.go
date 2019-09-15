package cmd

import (
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get User's Github Repositories and the issues",
	Long: `Get User's all Github repositories and the issues of the specify repository
gcli is a CLI tool for Github repositories and issues management, which will be improving continiously as the awesome tool!`,
	Run: func(cmd *cobra.Command, args []string) {
          Usage()
	},
}

func init() {
    //setCmd.Flags().StringP("issues", "i", "issues", "Get the issues of your specify repository!")
	rootCmd.AddCommand(getCmd)

}

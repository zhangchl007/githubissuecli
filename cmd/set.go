package cmd

import (
	"fmt"
	"github.com/zhangchl007/githubissuecli/pkg/github"
	"github.com/spf13/cobra"
)

// setCmd represents the set command
var userid, token, repo string
var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set User info for Github",
	Long: `Set User info including userid, user token and repo. gcli is a CLI tool for Github repositories and issues management,
which will be improving continiously as the awesome tool!`,
	Run: func(cmd *cobra.Command, args []string) {
         github.UpdateUserinfo(userid, token, repo)
         fmt.Println("The Userinfo config file had been changed successfully!")
	},
}


func init() {

    setCmd.Flags().StringVarP(&userid, "userid", "u", "zhang", "Change The Username!")
    setCmd.Flags().StringVarP(&token, "token", "t", "heeeee", "Change The User Token")
    setCmd.Flags().StringVarP(&repo, "repo", "r", "test", "Change The User Repo")
    rootCmd.AddCommand(setCmd)
}

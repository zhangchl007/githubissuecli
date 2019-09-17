package cmd

import (
	"fmt"
	"log"
	"github.com/zhangchl007/githubissuecli/pkg/github"
	"github.com/spf13/cobra"
)

// issuesCmd represents the issues command
var getissuesCmd = &cobra.Command{
	Use:   "issues",
    Short: "Get all the issues",
	Long: `Get User's all issues of the specify repository,gcli is a CLI tool for Github repositories and issues management,
which will be improving continiously as the awesome tool!`,
	Run: func(cmd *cobra.Command, args []string) {
        Userid, PersonalAccessToken, Repo := github.GetUserinfo()
        URL := github.IssuesURL + Userid + "/" + Repo + "/" + "issues?state=all"
        Action := github.MethodGet
        f, err := github.GetIssues(PersonalAccessToken, URL, Action)
        if err != nil {
		    log.Fatalln(err)
	    }
        fmt.Println("IssueNumber", "-----Createtime-----", "---Assignee---", "IssueState", "Lockstatus", "----IssueTitle----")
        for _, item := range *f {
            var userid string
            p := *item.Assignees
            for _, owner := range p {
                userid = owner.Login
            }
            fmt.Printf("#%-11d %10.20s %12s %10.55s %10t %20.20s\n", item.Number,item.CreateAt, userid, item.State, item.Locked, item.Title)
       }

	},
}
var createissueCmd = &cobra.Command{
    Use:   "issue",
    Short: "Create the issues for user repository",
    Long: `Create the issue for user repository, gcli is a CLI tool for Github repositories and issues management,
which will be improving continiously as the awesome tool`,
    //Args: cobra.MinimumNArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
      Filepath := "src/github.com/zhangchl007/githubissuecli/config/issue_template.yaml"
      PersonalAccessToken, url, Action, data := github.Parsehttpurl(Filepath, "open", "0", false)
      okay := github.UpdateIssues(PersonalAccessToken, url, Action, data)
      fmt.Println(okay)
    },
  }

var IssueNumber string
var closeissueCmd = &cobra.Command{
    Use:   "issue",
    Short: "Close the related issue for user repository",
    Long: `Close the related issue for user repository, gcli is a CLI tool for Github repositories and issues management,
which will be improving continiously as the awesome tool`,
    //Args: cobra.MinimumNArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
      Filepath := "src/github.com/zhangchl007/githubissuecli/config/issue_template.yaml"
      PersonalAccessToken, url, Action, data := github.Parsehttpurl(Filepath, "closed", IssueNumber, false)
      okay := github.UpdateIssues(PersonalAccessToken, url, Action, data)
      fmt.Println(okay)
    },
  }

var lockissueCmd = &cobra.Command{
    Use:   "issue",
    Short: "lock the issue for user repository",
    Long: `lock the issue for user repository, gcli is a CLI tool for Github repositories and issues management,
which will be improving continiously as the awesome tool`,
    //Args: cobra.MinimumNArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
      Filepath := "src/github.com/zhangchl007/githubissuecli/config/issue_template.yaml"
      PersonalAccessToken, url, Action, data := github.Parsehttpurl(Filepath, "null", IssueNumber, true)
      okay := github.UpdateIssues(PersonalAccessToken, url, Action, data)
      fmt.Println(okay)
    },
  }

var unlockissueCmd = &cobra.Command{
    Use:   "issue",
    Short: "unlock the issue for user repository",
    Long: `unlock the issue for user repository, gcli is a CLI tool for Github repositories and issues management,
which will be improving continiously as the awesome tool`,
    //Args: cobra.MinimumNArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
      Filepath := "src/github.com/zhangchl007/githubissuecli/config/issue_template.yaml"
      PersonalAccessToken, url, Action, data := github.Parsehttpurl(Filepath, "null", IssueNumber, false)
      okay := github.UpdateIssues(PersonalAccessToken, url, Action, data)
      fmt.Println(okay)
    },
  }

var updateissueCmd = &cobra.Command{
    Use:   "issue",
    Short: "update the issue for user repository",
    Long: `update the issue for user repository, gcli is a CLI tool for Github repositories and issues management,
which will be improving continiously as the awesome tool`,
    //Args: cobra.MinimumNArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
      Filepath := "src/github.com/zhangchl007/githubissuecli/config/issue_template.yaml"
      PersonalAccessToken, url, Action, data := github.Parsehttpurl(Filepath, "null", IssueNumber, false)
      okay := github.UpdateIssues(PersonalAccessToken, url, Action, data)
      fmt.Println(okay)
    },
  }

func init() {
	getCmd.AddCommand(getissuesCmd)
	createCmd.AddCommand(createissueCmd)
	closeCmd.AddCommand(closeissueCmd)
	lockCmd.AddCommand(lockissueCmd)
	unlockCmd.AddCommand(unlockissueCmd)
    updateCmd.AddCommand(updateissueCmd)
    closeissueCmd.Flags().StringVarP(&IssueNumber, "IssueNumber", "n", "0", "The issue Number")
    closeissueCmd.MarkFlagRequired("IssueNumber")
    lockissueCmd.Flags().StringVarP(&IssueNumber, "IssueNumber", "n", "0", "The issue Number")
    lockissueCmd.MarkFlagRequired("IssueNumber")
    unlockissueCmd.Flags().StringVarP(&IssueNumber, "IssueNumber", "n", "0", "The issue Number")
    unlockissueCmd.MarkFlagRequired("IssueNumber")
    updateissueCmd.Flags().StringVarP(&IssueNumber, "IssueNumber", "n", "0", "The issue Number")
    updateissueCmd.MarkFlagRequired("IssueNumber")
}

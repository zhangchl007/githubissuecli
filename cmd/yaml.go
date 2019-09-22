package cmd

import (
    "fmt"
    "github.com/zhangchl007/githubissuecli/pkg/github"
    "github.com/zhangchl007/githubissuecli/pkg/repos"
	"github.com/spf13/cobra"
)

// yamlCmd represents the yaml command
var issueyamlCmd = &cobra.Command{
	Use:   "issueyaml",
	Short: "Create the yaml file of issue template",
	Long:  `Create the yaml file of issue template, gcli is a CLI tool for Github repositories and issues management,
which will be improving continiously as the awesome tool!`,
	Run: func(cmd *cobra.Command, args []string) {
        var  yamlfile *github.Issueyamlfile
        myid := []string{Owner}
        mylabels := []string{Label}
        data, _, _:= yamlfile.UpdateIssueyaml(Title, Body, State, Locked, &myid, &mylabels)
        fmt.Printf("This is the converted joson file from issue template: %s\n", string(*data))
	},
}

var repoyamlCmd = &cobra.Command{
	Use:   "repoyaml",
	Short: "Create the yaml file of issue template",
	Long:  `Create the yaml file of issue template, gcli is a CLI tool for Github repositories and issues management,
which will be improving continiously as the awesome tool!`,
	Run: func(cmd *cobra.Command, args []string) {
        var  yamlfile *repos.Repoyamlfile
        data := yamlfile.UpdateRepoyaml(Name, Description, Homepage, Private)
        fmt.Printf("This is the converted joson file from repo template: %s\n", string(*data))
	},
}

func init() {
    issueyamlCmd.Flags().StringVarP(&Title, "Title", "t", "my issuexx", "The subject of your issue")
    issueyamlCmd.Flags().StringVarP(&Body, "Body", "b", "This is my issuexx", "The discription of your issue")
    issueyamlCmd.Flags().StringVarP(&State, "State", "s", "open", "The issue openning status")
    issueyamlCmd.Flags().BoolVarP(&Locked, "Locked", "l", false, "The issue locking status")
    issueyamlCmd.Flags().StringVarP(&Owner, "Owner", "u", "", "The issue owner")
    issueyamlCmd.Flags().StringVarP(&Label, "Label", "m", "bug", "The issue label")
    repoyamlCmd.Flags().StringVarP(&Name, "Name", "n", "test01", "The repo name for update soon")
    repoyamlCmd.Flags().StringVarP(&Description, "Description", "d", "this is my repo", "The repo description for update")
    repoyamlCmd.Flags().StringVarP(&Homepage, "Homepage", "w", "https://github.com/", "The homepage url of user's repo for update")
    repoyamlCmd.Flags().BoolVarP(&Private, "Private", "p", false, "The repo Public/Private setting for update")
	createCmd.AddCommand(issueyamlCmd)
	createCmd.AddCommand(repoyamlCmd)
}

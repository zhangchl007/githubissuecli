package cmd

import (
	"fmt"
	"github.com/zhangchl007/githubissuecli/pkg/github"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var Title, Body, State, Owner, Label string
var Locked bool

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create the yaml template of issue!",
	Long: `Create the yaml template of issue!, gcli is a CLI tool for Github repositories and issues management,
which will be improving continiously as the awesome tool!`,
	Run: func(cmd *cobra.Command, args []string) {
        var  yamlfile *github.Issueyamlfile
        myid := []string{Owner}
        mylabels := []string{Label}
        data, _, _:= yamlfile.UpdateIssueyaml(Title, Body, State, Locked, &myid, &mylabels)
        fmt.Println("This is the converted joson file from issue template: "+ string(*data))

	},
}

func init() {
    createCmd.Flags().StringVarP(&Title, "Title", "t", "Issue 7", "The subject of your issue")
    createCmd.Flags().StringVarP(&Body, "Body", "b", "This is my issue7", "The discription of your issue")
    createCmd.Flags().StringVarP(&State, "State", "s", "open", "The issue openning status")
    createCmd.Flags().BoolVarP(&Locked, "Locked", "l", false, "The issue locking status")
    createCmd.Flags().StringVarP(&Owner, "Owner", "u", "", "The issue owner")
    createCmd.Flags().StringVarP(&Label, "Label", "m", "bug", "The issue label")
	rootCmd.AddCommand(createCmd)

}

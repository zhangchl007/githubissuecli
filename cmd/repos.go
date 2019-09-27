package cmd

import (
	"fmt"
    "log"
	"github.com/zhangchl007/githubissuecli/pkg/repos"
	"github.com/spf13/cobra"
)

// repoCmd represents the repo command
var getreposCmd = &cobra.Command{
	Use:   "repos",
	Short: "Get user repositories",
	Long: `Get user repositories, gcli is a CLI tool for Github repositories and issues management,
which will be improving continiously as the awesome tool!`,
	Run: func(cmd *cobra.Command, args []string) {
        PersonalAccessToken, URL, Action, _ := repos.Parsehttpurl("GET","null")
        f, err := repos.GetRepos(PersonalAccessToken, URL, Action)
        if err != nil {
		    log.Fatalln(err)
	    }
        fmt.Println("-------------Reponame-------------", "-Createtime-", "--Owner--", "-Private-", "-Openissues-", "---------Description-----------")
        for _, item := range *f {
            fmt.Printf("%-35s %5.10s %12s %t %5d %-12s\n", item.Name, item.CreateAt, item.Owner.Login, item.Private, item.Openissues, item.Description)
        }

	},
}

var createrepoCmd = &cobra.Command{
    Use:   "repo",
    Short: "Create the repository for the User",
    Long: `Create the repository for the User, gcli is a CLI tool for Github repositories and issues management,
which will be improving continiously as the awesome tool`,
    //Args: cobra.MinimumNArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        //Filepath := "src/github.com/zhangchl007/githubissuecli/config/repo_template.yaml"
        PersonalAccessToken, URL, Action, Data := repos.Parsehttpurl("POST", Name)
        okay := repos.UpdateRepos(PersonalAccessToken, URL, Action, Data)
        fmt.Println(okay)
    },
  }

var updaterepoCmd = &cobra.Command{
    Use:   "repo",
    Short: "Create the repository for the User",
    Long: `Create the repository for the User, gcli is a CLI tool for Github repositories and issues management,
which will be improving continiously as the awesome tool`,
    //Args: cobra.MinimumNArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        var  yamlfile *repos.Repoyamlfile
        PersonalAccessToken, URL, Action, _ := repos.Parsehttpurl("PATCH", Name)
        Data := yamlfile.UpdateRepoyaml(Name, Description, Homepage, Private)
        okay := repos.UpdateRepos(PersonalAccessToken, URL, Action, Data)
        fmt.Println(okay)
    },
  }

var lockrepoCmd = &cobra.Command{
    Use:   "repo",
    Short: "Create the repository for the User",
    Long: `Create the repository for the User, gcli is a CLI tool for Github repositories and issues management,
which will be improving continiously as the awesome tool`,
    //Args: cobra.MinimumNArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        PersonalAccessToken, URL, Action, Data := repos.Parsehttpurl("lock", Name)
        okay := repos.UpdateRepos(PersonalAccessToken, URL, Action, Data)
        fmt.Println(okay)
    },
  }

var unlockrepoCmd = &cobra.Command{
    Use:   "repo",
    Short: "Create the repository for the User",
    Long: `Create the repository for the User, gcli is a CLI tool for Github repositories and issues management,
which will be improving continiously as the awesome tool`,
    //Args: cobra.MinimumNArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        PersonalAccessToken, URL, Action, Data := repos.Parsehttpurl("unlock", Name)
        okay := repos.UpdateRepos(PersonalAccessToken, URL, Action, Data)
        fmt.Println(okay)
    },
  }

var deleterepoCmd = &cobra.Command{
    Use:   "repo",
    Short: "Create the repository for the User",
    Long: `Create the repository for the User, gcli is a CLI tool for Github repositories and issues management,
which will be improving continiously as the awesome tool`,
    //Args: cobra.MinimumNArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        PersonalAccessToken, URL, Action, Data := repos.Parsehttpurl("DELETE", Name)
        okay := repos.UpdateRepos(PersonalAccessToken, URL, Action, Data)
        fmt.Println(okay)
    },
  }

func init() {
    createrepoCmd.Flags().StringVarP(&Name, "Name", "n", "test01", "The repo name for creation soon")
    updaterepoCmd.Flags().StringVarP(&Name, "Name", "n", "test01", "The repo name for info update")
    lockrepoCmd.Flags().StringVarP(&Name, "Name", "n", "test01", "The specify repo name for locking")
    unlockrepoCmd.Flags().StringVarP(&Name, "Name", "n", "test01", "The specify repo name for unlocking")
    deleterepoCmd.Flags().StringVarP(&Name, "Name", "n", "test01", "The specify repo name for removing")
    updaterepoCmd.Flags().StringVarP(&Description, "Description", "d", "this is my repo update", "The repo description for update")
    updaterepoCmd.Flags().StringVarP(&Homepage, "Homepage", "w", "https://github.com/update",  "The repo homepage url for update")
    updaterepoCmd.Flags().BoolVarP(&Private, "Private", "l", false,  "The parameter of public/private repo setting")
    createrepoCmd.MarkFlagRequired("Name")
    updaterepoCmd.MarkFlagRequired("Name")
    lockrepoCmd.MarkFlagRequired("Name")
    unlockrepoCmd.MarkFlagRequired("Name")
    deleterepoCmd.MarkFlagRequired("Name")
	createCmd.AddCommand(createrepoCmd)
	updateCmd.AddCommand(updaterepoCmd)
    lockCmd.AddCommand(lockrepoCmd)
    unlockCmd.AddCommand(unlockrepoCmd)
    deleteCmd.AddCommand(deleterepoCmd)
	getCmd.AddCommand(getreposCmd)
}

package repos
import (
    "fmt"
    "regexp"
    "github.com/zhangchl007/githubissuecli/pkg/github"
    "testing"
)

func TestParsehttpurl(t *testing.T) {
    var URL string
    var Data []byte
    Repo := "test01"
    s := [...]string{"GET", "POST", "PUT", "DELETE", "PATCH", "HEAD"}
    Userid, PersonalAccessToken, _ := github.GetUserinfo()

    for _, Action := range s {
        if Action == github.MethodGet {
            URL = github.ReposURL + "?per_page=100"
        } else if Action == github.MethodPost {
            URL = github.ReposURL
            Data = []byte(`{"name": "{{.Name}}", "auto_init": false, "gitignore_template": "nanoc"}`)
            //Data =[]byte(`{"name": "{{.Name}}", "auto_init": false}`)
            re, _ := regexp.Compile("{{.Name}}")
            Data = re.ReplaceAll(Data, []byte(Repo))
        } else if Action == github.MethodPatch {
            URL = github.UserReposURL + Userid + "/" + Repo
        } else if Action == "lock" {
            URL = github.UserReposURL + Userid + "/" + Repo
            Action = github.MethodPatch
            Data = []byte(`{"private": true}`)
        } else if Action == "unlock" {
            URL = github.UserReposURL + Userid + "/" + Repo
            Action = github.MethodPatch
            Data = []byte(`{"private": false}`)
        } else if Action == github.MethodDelete {
            URL = github.UserReposURL + Userid + "/" + Repo
        } else {
            fmt.Println("Please input the right parameters!")
        }
        fmt.Println(PersonalAccessToken, URL, Action, &Data)
    }
}


package github

import (
    "fmt"
    "os"
    "testing"
)

func TestParsehttpurl(t *testing.T) {
    var url string
    var data []byte
    GOPATH := os.Getenv("GOPATH")
    Filepath := GOPATH + "/src/github.com/zhangchl007/githubissuecli/config/issue_template.yaml"
    IssueNumber := "100"
    s := [...]string{"GET", "POST", "PUT", "DELETE", "PATCH", "HEAD"}
    //data := []byte(`{"assignees":[],"body":"speed up golang development","labels":["bug"],"locked":false,"state":"open","title":"this issue7"}`)
    Userid, PersonalAccessToken, Repo := GetUserinfo()
    for _, Action := range s {
        State := "open"
        Locked := false
        if State == "open" && Locked == false {
            url = IssuesURL + Userid + "/" + Repo + "/issues"
            Action = MethodPost
            data = ReadTemplate(Filepath)
        } else if Filepath == "" {
            url = IssuesURL + Userid + "/" + Repo + "/issues/" + IssueNumber
            Action = MethodPatch
        } else if State == "closed" && len(IssueNumber) > 0 {
            url = IssuesURL + Userid + "/" + Repo + "/issues/" + IssueNumber
            Action = MethodPatch
            data = []byte(`{"state":"close"}`)
        } else if Locked && len(IssueNumber) > 0 {
            url = IssuesURL + Userid + "/" + Repo + "/issues/" + IssueNumber + "/lock"
            Action = MethodPut
            data = []byte(`{"locked":true}`)
        } else if Locked == false && len(IssueNumber) > 0 {
            url = IssuesURL + Userid + "/" + Repo + "/issues/" + IssueNumber + "/lock"
            Action = MethodDelete
        } else {
            fmt.Println("Please input the right parameters!")
        }
       fmt.Println(PersonalAccessToken, url, Action, &data)
    }

}

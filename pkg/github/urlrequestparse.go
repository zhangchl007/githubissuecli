package github
import (
    "fmt"
)

func Parsehttpurl(Filepath, State, IssueNumber string, Locked bool) (string, string, string, *[]byte) {
    var url, Action string
    var data []byte
    Userid, PersonalAccessToken, Repo := GetUserinfo()
    if State =="open" && Locked == false {
        url = IssuesURL + Userid + "/" + Repo + "/issues"
        fmt.Println(url)
        Action = MethodPost
        data = ReadTemplate(Filepath)
    } else if Filepath == "" {
        url = IssuesURL + Userid + "/" + Repo + "/issues/" + IssueNumber
        Action = MethodPatch
    } else if State == "closed" && len(IssueNumber) > 0 {
        url = IssuesURL + Userid + "/" + Repo + "/issues/" + IssueNumber
        Action = MethodPatch
        data = []byte(`{"state":"close"}`)
    } else if Locked  && len(IssueNumber) > 0 {
        url = IssuesURL + Userid + "/" + Repo + "/issues/" + IssueNumber + "/lock"
        Action = MethodPut
        data = []byte(`{"locked":true}`)
    } else if Locked == false && len(IssueNumber) > 0  {
        url = IssuesURL + Userid + "/" + Repo + "/issues/" + IssueNumber + "/lock"
        Action = MethodDelete
    } else {
        fmt.Println("Please input the right parameters!")
    }

    return PersonalAccessToken, url, Action, &data
}

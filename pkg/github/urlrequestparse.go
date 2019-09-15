package github
import (
    "fmt"
)

func Parsehttpurl(Filepath, State, IssueNumber string, Locked bool) (string, string, string, *[]byte) {
    var url, Action string
    data := ReadTemplate(Filepath)
    Userid, PersonalAccessToken, Repo := GetUserinfo()
    if State =="open" && Locked == false {
        url = IssuesURL + Userid + "/" + Repo + "/issues"
        Action = MethodPost
    } else if State == "closed" && len(IssueNumber) > 0 {
        url = IssuesURL + Userid + "/" + Repo + "/issues/" + IssueNumber
        Action = MethodPatch
        fmt.Println(url)
    } else if Locked  && len(IssueNumber) > 0 {
        url = IssuesURL + Userid + "/" + Repo + "/issues/" + IssueNumber + "/lock"
        Action = MethodPut
    } else {
        fmt.Println("Please input the right parameters!")
    }

    return PersonalAccessToken, url, Action, &data
}

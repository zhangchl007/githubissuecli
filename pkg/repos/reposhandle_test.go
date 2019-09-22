package repos
import (
    "encoding/json"
    "github.com/zhangchl007/githubissuecli/pkg/github"
    "bytes"
    "fmt"
    "log"
    "time"
    "net/http"
)

func GetRepos(PersonalAccessToken, URL, Action string) (*[]github.Repos, error) {
    req, err := http.NewRequest(Action, URL,  nil)
    if err != nil {
        log.Fatal("Error reading request.")
    }
    req.Header.Add("Authorization", "Bearer " + PersonalAccessToken)
    //req.Header.Add("User-Agent","Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/50.0.2661.102 Safari/537.36")
    //req.Header.Add("Content-Type","application/json")
    client := &http.Client{}
    //Send req using http Client
    resp, err := client.Do(req)
    if err != nil {
        return nil, err
    }

    if resp.StatusCode != http.StatusOK {
        resp.Body.Close()
        return nil, fmt.Errorf("search query failed: %s", resp.Status)
    }

    var result []github.Repos

    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        resp.Body.Close()
        return nil, err
    }

    resp.Body.Close()
    return &result, nil
}

func UpdateRepos(PersonalAccessToken,  URL, Action string, data *[]byte)(string) {
    req, err := http.NewRequest(Action, URL,  bytes.NewBuffer(*data))
    if err != nil {
        log.Fatal("Error reading request.")
    }
    req.Header.Add("Content-Type","application/json")
    req.Header.Add("Authorization", "Bearer " + PersonalAccessToken)
    // Create and Add cookie to request
	//cookie := http.Cookie{Name: "cookie_name", Value: "cookie_value"}
	//req.AddCookie(&cookie)
	// Set client timeout
	client := &http.Client{Timeout: time.Second * 3}
    //Send req using http Client
    resp, err := client.Do(req)
    if err != nil {
        fmt.Printf("err: %v\n", err)
    }

    if resp.StatusCode != http.StatusOK {
        resp.Body.Close()
        fmt.Errorf("The repo is failed to create!: %s", resp.Status)
    }
    defer resp.Body.Close()
    //fmt.Println("response Status:", resp.Status)
	//fmt.Println("response Headers:", resp.Header)
    return resp.Status
}

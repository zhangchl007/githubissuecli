package github
import (
//    "flag"
    "encoding/json"
//    "strings"
    "bytes"
    "fmt"
    "os"
    "io/ioutil"
    "log"
    "time"
    "net/http"
    y2j "github.com/ghodss/yaml"
    "github.com/spf13/viper"
    yaml "gopkg.in/yaml.v2"
)

func GetIssues(PersonalAccessToken, URL, Action string) (*[]Issues, error) {
    //url := IssuesURL + Userid + "/" + Repo + "/" + "issues?state=all"
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

    var result []Issues

    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        resp.Body.Close()
        return nil, err
    }

    resp.Body.Close()
    return &result, nil
}

func (yamlfile *Issueyamlfile) UpdateIssueyaml(Title, Body, State string, Locked bool, Assignees, Labels *[]string) (*[]byte, string, bool){
    tmpfile := "/tmp/a.txt"
    IssueTemplate :="issue_template"
    IssueyamlPath  :="src/github.com/zhangchl007/githubissuecli/config/"
    TemplateFile := IssueyamlPath + IssueTemplate + ".yaml"
    viper.SetConfigName(IssueTemplate)
    viper.AddConfigPath(IssueyamlPath)
    err := viper.ReadInConfig()
    if err != nil {
        log.Fatal(err)
    }

    err = viper.Unmarshal(&yamlfile)
    if err != nil{
       log.Fatalf("unable to decode into struct, %v", err)
    }
   //viper.Set(yamlfile.Title, Title)
   yamlfile.Title = Title
   yamlfile.Body  = Body
   yamlfile.State = State
   yamlfile.Locked = Locked
   yamlfile.Assignees = *Assignees
   yamlfile.Labels = *Labels

   //viper.WriteConfig()

   // encode yaml again
   d, err := yaml.Marshal(&yamlfile)
   if err != nil {
      log.Fatalf("error: %v", err)
   }
   WriteToFile(tmpfile, d)
   MoveFile(tmpfile, TemplateFile)
   fmt.Printf("The yaml file of issue template: %s had been created/updated succesfully!\n",TemplateFile)
   //yaml to json for issue creation
   y2 := []byte(string(d))
   j2, err := y2j.YAMLToJSON(y2)
   if err != nil {
       fmt.Printf("err: %v\n", err)
    }
   //fmt.Println(string(j2))

   return &j2, State, Locked
}

// generate the tempfile for yamlfile
func WriteToFile(f string ,d []byte)  {
    err := ioutil.WriteFile(f, d, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

//replace yamlfile file
func MoveFile(from,to string) {
	err := os.Rename(from, to)
	if err != nil {
		log.Fatal(err)
	}
}

func UpdateIssues(PersonalAccessToken,  URL, Action string, data *[]byte)(string) {
    //url := IssuesURL + Userid + "/" + Repo + "/issues"
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
	client := &http.Client{Timeout: time.Second * 10}
    //Send req using http Client
    resp, err := client.Do(req)
    if err != nil {
        fmt.Printf("err: %v\n", err)
    }

    if resp.StatusCode != http.StatusOK {
        resp.Body.Close()
        fmt.Errorf("The issue is failed to create!: %s", resp.Status)
    }
    defer resp.Body.Close()
    //fmt.Println("response Status:", resp.Status)
	//fmt.Println("response Headers:", resp.Header)
    return resp.Status

}


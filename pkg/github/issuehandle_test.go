package github
import (
    "encoding/json"
    "bytes"
    "fmt"
    "os"
    "io/ioutil"
    "log"
    "time"
    "net/http"
    "testing"
    y2j "github.com/ghodss/yaml"
    "github.com/spf13/viper"
    yaml "gopkg.in/yaml.v2"
)

func TestGetIssues(t *testing.T) {
    PersonalAccessToken := "xxxxxxxxxxxxxxxx"
    URL := "https://api.github.com/repos/zhangchl007/test/issues"
    Action := "GET"
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
        log.Fatal(err)
    }

    if resp.StatusCode != http.StatusOK {
        resp.Body.Close()
        fmt.Errorf("search query failed: %s", resp.Status)
    }

    //var result []Issues
     result := []struct {
         Number int
         HTMLURL string
         Title string
         State string
         Locked bool
         Assignees *[]Assignees
         CreateAt time.Time
         Body string
     }{
         {
             Number: 2,
             HTMLURL: "https://github.com/zhangchl007/test/issues/2",
             Title: "this is issue8",
             State: "closed",
             Locked: false,
             Assignees: &[]Assignees{
                 {
                     Login: "zhangchl007",
                     HTMLURL: "https://github.com/zhangchl007/test",
                 },
             },
         CreateAt: time.Now(),
         Body:     "I'm having a problem with this",
        },
     }
    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        resp.Body.Close()
    }

    resp.Body.Close()
    fmt.Println (&result)
}

func TestUpdateIssueyaml(t *testing.T){

    yamlfile := []struct {
        Title     string
        Body      string
        Assignees []string
        State     string
        Locked    bool
        Labels    []string
    } {
        {
            Title: "this is issue8",
            Body:  "this is my issue, Please take care of it",
            Assignees: []string{ "zhangchl007" },
            State: "open",
            Locked: false,
            Labels: []string{ "bug" },
        },
    }
    tmpfile := "/tmp/a.test"
    IssueTemplate :="issue_template"
    IssueyamlPath  :="../../config/"
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
   //yamlfile.Title = Title
   //yamlfile.Body  = Body
   //yamlfile.State = State
   //yamlfile.Locked = Locked
   //yamlfile.Assignees = *Assignees
   //yamlfile.Labels = *Labels

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

   fmt.Println(&j2)
}

// generate the tempfile for yamlfile
func TestWriteToFile(t *testing.T)  {
    f := "/tmp/a.test"
    d := []byte(`{"kind":"Test","apiVersion":"other/blah","spec": {"A": 1, "b": 2, "h": 3, "I": 4}}`)
    err := ioutil.WriteFile(f, d, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

//replace yamlfile file
func TestMoveFile(t *testing.T) {
    from := "/tmp/a.test"
    to   := "/tmp/b.test"
	err := os.Rename(from, to)
	if err != nil {
		log.Fatal(err)
	}
}

func TestUpdateIssues(t *testing.T) {

    s := [...]string{"GET", "POST", "PUT", "DELETE", "PATCH", "HEAD"}
    for _, Action := range s {
        PersonalAccessToken := "xxxxxxxxxxxxxxxxxxx"
        URL := "https://api.github.com/repos/zhangchl007/test/issues"
        data := []byte(`{"assignees":[],"body":"speed up golang development","labels":["bug"],"locked":false,"state":"open","title":"this issue7"}`)
        req, err := http.NewRequest(Action, URL, bytes.NewBuffer(data))
        if err != nil {
            log.Fatal("Error reading request.")
        }

        req.Header.Add("Content-Type", "application/json")
        req.Header.Add("Authorization", "Bearer "+PersonalAccessToken)
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
        fmt.Println(resp.Status)
    }
}

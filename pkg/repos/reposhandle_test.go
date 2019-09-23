package repos
import (
    "encoding/json"
    "github.com/zhangchl007/githubissuecli/pkg/github"
    "os"
    "io/ioutil"
    "bytes"
    "fmt"
    "testing"
    "time"
    "log"
    "net/http"
    y2j "github.com/ghodss/yaml"
    yaml "gopkg.in/yaml.v2"

)

func TestGetRepos(t *testing.T) {
    _, PersonalAccessToken, _ := github.GetUserinfo()
    URL := "https://api.github.com/repos/zhangchl007/test/issues"
    Action := "GET"
    req, err := http.NewRequest(Action, URL,  nil)
    if err != nil {
        log.Fatal("Error reading request.")
    }
    req.Header.Add("Authorization", "Bearer " + PersonalAccessToken)
    //req.Header.Add("User-Agent","Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/50.0.2661.102 Safari/537.36")
    req.Header.Add("Content-Type","application/json")
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

    //var result []github.Repos
    result :=[]struct{
        ID   int
        Name string
        Private bool
        Owner   github.Owner
        HTMLURL string
        Description string
        Homepage string
        Fork  bool
        Teamurl string
        CreateAt time.Time
        Forks int
        Openissues int
        Language string
    }{
        {
            ID: 100,
            Name: "test01",
            Private: false,
            Owner: github.Owner{
                Login: "zhangchl007",
            },
            Description: "this is my repo",
            Homepage: "https://github.io/zhangchl007",
            Fork: false,
            Teamurl: "https://github.io/zhangchl007",
            CreateAt: time.Now(),
            Forks: 20,
            Openissues: 14,
            Language:   "golang",
        },
    }

    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        resp.Body.Close()
        fmt.Errorf("search query failed: %s", resp.Status)
    }

    resp.Body.Close()
     fmt.Println(&result)
}

func (yamlfile *Repoyamlfile) TestUpdateRepoyaml(t *testing.T){
    tmpfile := "/tmp/b.test"
    RepoTemplate := "repo_template.yaml"
    GOPATH := os.Getenv("GOPATH")
    RepoyamlPath  := GOPATH + "/src/github.com/zhangchl007/githubissuecli/config/"
    TemplateFile := RepoyamlPath + RepoTemplate
    content, err := ioutil.ReadFile(TemplateFile)
    if err != nil {
        log.Fatal(err)
    }
   // struc Unmasrshal
   err = yaml.Unmarshal(content, &yamlfile)
   if err != nil {
     panic(err)
   }
   yamlfile.Name = "test01"
   yamlfile.Description  = "This is my repo"
   yamlfile.Homepage = "https://homepage.com"
   yamlfile.Private = false

   // encode yaml again
   d, err := yaml.Marshal(&yamlfile)
   if err != nil {
      log.Fatalf("error: %v", err)
   }
   WriteToFile(tmpfile, d)
   MoveFile(tmpfile, TemplateFile)
   fmt.Printf("The yaml file of repo template: %s had been created/updated succesfully!\n",TemplateFile)
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

func TestUpdateRepos(t *testing.T) {
    _, PersonalAccessToken, _ := github.GetUserinfo()
    s := [...]string{"GET", "POST", "PUT", "DELETE", "PATCH"}
    URL := "https://api.github.com/repos/zhangchl007/test/issues"
    data := []byte(`{"name": "test01","description": "this is my repo update","homepage": "https://github.com/","private": false}`)
    for _, Action := range s {
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
        client := &http.Client{Timeout: time.Second * 20}
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
        fmt.Println(resp.Status)
    }
}

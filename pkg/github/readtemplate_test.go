package github

import (
	"fmt"
	y2j "github.com/ghodss/yaml"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestReadTemplate(t *testing.T) {
    GOPATH := os.Getenv("GOPATH")
	Filepath := GOPATH + "/src/github.com/zhangchl007/githubissuecli/config/issue_template.yaml"
	content, err := ioutil.ReadFile(Filepath)
	if err != nil {
		log.Fatal(err)
	}

    j2, err := y2j.YAMLToJSON(content)
    if err != nil {
       fmt.Printf("err: %v\n", err)
    }
	fmt.Printf("The post json contents: %s\n", j2)
}


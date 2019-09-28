package github

import (
	"fmt"
	"io/ioutil"
	"log"

	y2j "github.com/ghodss/yaml"
)

// read the yaml file of repo/issue template, which will be removed in the future
func ReadTemplate(Filepath string) []byte {
	content, err := ioutil.ReadFile(Filepath)
	if err != nil {
		log.Fatal(err)
	}

	j2, err := y2j.YAMLToJSON(content)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Println("The post json contents")
	return j2
}

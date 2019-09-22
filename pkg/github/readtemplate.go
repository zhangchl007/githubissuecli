package github
import (
     "fmt"
     "log"
     y2j "github.com/ghodss/yaml"
     "io/ioutil"
)

func ReadTemplate(Filepath string)([]byte) {
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


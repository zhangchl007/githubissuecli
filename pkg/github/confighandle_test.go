package github
import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"os"
	"testing"
)

func TestUpdateUserinfo(t *testing.T) {
	id := "zhangchl007"
	idtoken := "xxxxxxxxxxxxx"
	idrepo := "test"
	GOPATH := os.Getenv("GOPATH")
	ConfigPath  := GOPATH + "/src/github.com/zhangchl007/githubissuecli/config"

    ConfigTemplate :="test"
    //ConfigPath  :="../../config"

    viper.SetConfigName(ConfigTemplate)
    viper.AddConfigPath(ConfigPath)
    err := viper.ReadInConfig()
    if err != nil {
        log.Fatal(err)
    }

   viper.Set("userid", id)
   viper.Set("token", idtoken)
   viper.Set("repo", idrepo)
   viper.WriteConfig()
   viper.WatchConfig()
   viper.OnConfigChange(func(e fsnotify.Event) {
   fmt.Println("配置发生变更：", e.Name)
   })
}

func TestGetUserinfo(t *testing.T) {
	GOPATH := os.Getenv("GOPATH")
    ConfigTemplate :="config"
    ConfigPath  := GOPATH + "/src/github.com/zhangchl007/githubissuecli/config"
    viper.SetConfigName(ConfigTemplate)
    viper.AddConfigPath(ConfigPath)
    err := viper.ReadInConfig()
    if err != nil {
        log.Fatal(err)
    }
    userid := viper.GetString("userid")
    token := viper.GetString("token")
    repo := viper.GetString("repo")
    fmt.Println(userid, token, repo)
}


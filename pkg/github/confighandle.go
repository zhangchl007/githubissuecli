package github

import (
	"fmt"
	"log"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

//update user info as the parameter
func UpdateUserinfo(id, idtoken, idrepo string) {
	ConfigTemplate := "config"
	GOPATH := os.Getenv("GOPATH")
	ConfigPath := GOPATH + "/src/github.com/zhangchl007/githubissuecli/config/"
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
	fmt.Printf("The Userinfo config file: %s had been changed successfully!\n", ConfigPath+ConfigTemplate+".yaml")
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("配置发生变更：", e.Name)
	})
}

// Get user info including userid/token/passwd
func GetUserinfo() (string, string, string) {
	GOPATH := os.Getenv("GOPATH")
	ConfigTemplate := "config"
	ConfigPath := GOPATH + "/src/github.com/zhangchl007/githubissuecli/config/"
	viper.SetConfigName(ConfigTemplate)
	viper.AddConfigPath(ConfigPath)
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	userid := viper.GetString("userid")
	token := viper.GetString("token")
	repo := viper.GetString("repo")
	return userid, token, repo
}

package github
import (
    "fmt"
    "log"
    "github.com/fsnotify/fsnotify"
    "github.com/spf13/viper"
)

func UpdateUserinfo(id, idtoken, idrepo string) {
    ConfigTemplate :="config"
    ConfigPath  :="src/github.com/zhangchl007/githubissuecli/config/"
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

func GetUserinfo()(string,string,string) {
//    var s []string
    ConfigTemplate :="config"
    ConfigPath  :="src/github.com/zhangchl007/githubissuecli/config/"
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


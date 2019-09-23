# `gcli` is a CLI tool for Github repositories and issues management,which will be improving continiously as the awesome tool!


# Quick Start

Download source code:
```bash
# go get -u github.com/zhangchl007/githubissuecli

# go build -o bin/gcli  src/github.com/zhangchl007/githubissuecli/main.go

# ls -l $GOPATH/bin/gcli
-rwxr-xr-x 1 jimmy szadmin 12498549 9月  23 07:19 /home/jimmy/Downloads/go/bin/gcli
```
# Configuration

You must generate the github token for personal github repo and issue management.

The yaml file of User authentication for github
```
#cat src/github.com/zhangchl007/githubissuecli/config/config.yaml
repo: test
token: xxxxxxxxxxxxx
userid: zhangchl007
```

The yaml file of repo template
```
#cat src/github.com/zhangchl007/githubissuecli/config/repo_template.yaml 
 name: test01
 description: this is my repo update
 homepage: https://github.com/
 private: false

```
# Basic Commands

```
#Set userinfo
gcli set  -u username -t token -r repo

#create the yaml file of repo template
gcli create repoyaml -n "test01" -d "this is my repo, take care of it"

#create the yaml file of issue template
gcli create issueyaml --Title "issue8" --Body "this is my issue8, take care of it"

#########repo and issue management################
#get user repo and the issues of repo
gcli get repos
gcli get issues
#create repo and issue
gcli create repo -n "test01"
gcli create issue

#update repo and issue
gcli update repo -n test01 -d "this is my repo update, take care of it"
gcli update issue -n issuenumber --Body "this is my issue8, take care of it"

#issue locked and public/private repo setting
gcli lock repo -n  test01
gcli lock issue -n issuenumber

#Issue unlocked and public/private repo setting
gcli lock repo -n  test01
gcli lock issue -n issuenumber

#close issue
gcli close issue -n issuenumber

#delete repo
gcli delete repo -n test01

```
# Roadmap

`gcli` is a CLI tool for Github repositories and issues management, which will be improving continiously as the awesome tool!

- [ ] Repo and Issue Batch managment: currently, `gcli` only provide the basic operation of github repo and issue.
- [ ] Multi authentication way of github: currently, `gcli` is only sample authentication way by github token.
- [ ] More funture extension
- [ ] e2e tests

If you are interested in any of the above features, please file an issue to avoid potential duplication.

# Contribute

Feel free to open issues and pull requests. Any feedback is highly appreciated!

# Acknowledgement

This project because of learning cobra framwork with kubernetes



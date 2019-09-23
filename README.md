#gcli is a CLI tool for Github repositories and issues management,
which will be improving continiously as the awesome tool!

```
# go get -u github.com/zhangchl007/githubissuecli

# go build -o bin/gcli  src/github.com/zhangchl007/githubissuecli/main.go

# ls -l $GOPATH/bin/gcli
-rwxr-xr-x 1 jimmy szadmin 12498549 9月  23 07:19 /home/jimmy/Downloads/go/bin/gcli

#Set userinfo
gcli set  -u username -t token -r repo

#create the yaml file of repo template
gcli create repoyaml -n "test01" -d "this is my repo, take care of it"

#create the yaml file of issue template
gcli create issueyaml --Title "issue8" --Body "this is my issue8, take care of it"

#########repo and issue management################
#get user repo and the issues of repo
gcli get repos/issues

#create repo and issue
gcli create repo -n "test01"/issue
gcli create issue

#update repo and issue
gcli update repo -n test01 -d "this is my repo update, take care of it"
gcli update issue -n issuenumber --Body "this is my issue8, take care of it"

#issue locked and public/private repo
gcli lock repo -n  test01
gcli lock issue -n issuenumber

#close issue
gcli close issue -n issuenumber

#delete repo
gcli delete repo -n test01

```


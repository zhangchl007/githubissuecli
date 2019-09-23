package cmd
import (
   "fmt"
)

func Usage() {
     message := `#Set userinfo
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
gcli delete repo -n test01`

fmt.Printf("%s\n",message)
}




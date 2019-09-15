# githubissuecli
gcli is a CLI tool for Github repositories and issues management,
which will be improving continiously as the awesome tool!
#Set userinfo
gcli set  -u username -t token -r repo

#create issue yaml template
gcli create --Title "issue8" --Body "this is my issue8, take care of it"

#repo and issue management
#get user repo and the issues of repo

gcli get issue /repo
 
#create repo and issue
gcli create issue/repo

#update repo and issue
gcli update issue/repo -n issuenumber

#issue locked and public/private repo
gcli lock issue/repo -n issuenumber 

#close issue
gcli close issue -n issuenumber

#assign the issue to owner
gcli assign issuenumber owner 

#delete repo

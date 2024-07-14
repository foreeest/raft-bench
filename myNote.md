# develop note #

## run ##
can not go get dragonboat/v3, but can copy from previous go.mod and go.sum  
then 
```shell
$ go get github.com/lni/dragonboat/v3@latest
$ go get github.com/lni/dragonboat/v4@master
$ go build
$ goreman -f Procfile-dragonboat start

```
dont run`go mod tidy`, I dont understand  
change the dragonboat v4 into v3  
modify the name

can not underatand  
`git config --global --add url."git@github.com".insteadOf "https://github.com/"`  
then I can't push  
`git config --global --add url."https://github.com/".insteadOf "git@github.com"`  
doing this is useless  
`git config --global --unset url."git@github.com:".insteadOf`   
`git config --global --list`  

well I need to
`go mod tidy github.com/foreeest/raftbench`

## using Prometheus ##


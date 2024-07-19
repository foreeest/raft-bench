# develop note #

## run ##
change the head file of dragonboat from v4 into v3 
modify the head file name into foreeest
```shell
$ go mod init github.com/foreeest/raftbench
$ go mod tidy
$ go build
$ goreman -f Procfile-dragonboat start
```

- some exp:
can not underatand  
`git config --global --add url."git@github.com".insteadOf "https://github.com/"`  
then I can't push  
`git config --global --add url."https://github.com/".insteadOf "git@github.com"`  
doing this is useless  
`git config --global --unset url."git@github.com:".insteadOf`   
`git config --global --list`  
well I need to
`go mod init github.com/foreeest/raftbench` 

go can not import each others?

## using Prometheus ##

why is that it can't run sometimes  
`ps -elf|grep raftbench`  
I think maybe control c have problem  

dont forget to `go build`   
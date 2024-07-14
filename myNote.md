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

## using Prometheus ##


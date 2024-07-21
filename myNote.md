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
doing this is useless, this err is because of wrong url of lib    
`git config --global --unset url."git@github.com:".insteadOf`   
`git config --global --list`  
well I need to
`go mod init github.com/foreeest/raftbench` 

go can not import each others?

## using Prometheus ##

why is that it can't run sometimes?    
`ps -elf|grep raftbench`  
I think maybe control c have problem  

**don't forget** to `go build`,everytime you modify the file   

## running v1.0.0 ##

is `busy false` right? why the `lni v4` version less busy false?what is the difference?    
the `step` param is too small?  
why is the last has more failure?     

### how to switch between `lni v4` and `foreeest udp` ###   
when modifying the header, just modify this, don't modify `foreeest/raftbench`   

- **udp -> v4**  

```shell
$ rm -r go.mod go.sum
$ go mod init github.com/foreeest/raftbench
# then modify the header in code
# add replace in go.mod `replace github.com/foreeest/raftbench => ../raftbench`  
$ go mod tidy
$ go build
$ rm -r wal-*
```
but on github.com `raftbench` is not update, so this is **useless**, you need to use replace     
actually this is a little confusing, I don't need to **replace** dragonboat; but I need to **replace** raftbench to local  
Note that sometimes,not adding **replace** is right,but adding it you never wrong  

```shell
$ goreman -f Procfile-drgonboat start
```

- **v4 -> replace**   
almost the same as udp->v4    

## get better performance ##

- param
step set higher?   
- why modifying the first send will get connection refused?  


## sending snapshot? ##

what is the sz of snapshot?  

why that tcp is sticked?  
# Setting Environment
---
First GOPATH will be seen as a library location for “go get”  
All the rest will be seen as a workspace by adding /src  

For exmaple:
```
export GOPATH=~/Documents/go/golib
export PATH=$PATH:$GOPATH/bin
export GOPATH=$GOPATH:~/Documents/go/GatorStore/backend
```

Workspace path is down below: 
```
~/Documents/go/GatorStore/backend/src
```

Binary file will is in the path below:
```
~/Documents/go/GatorStore/backend/bin
```

Binary package be linked is in the path below:
```
~/Documents/go/GatorStore/backend/pkg
```

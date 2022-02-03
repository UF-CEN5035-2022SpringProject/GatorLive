# Golang Version 
---
**go1.17.6**

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
# Download Dependencies
---
```
go get -u -v -f all
```

# Run Go program
---
Differences between the three cmd below:
[go run vs go build vs go install](https://levelup.gitconnected.com/go-run-vs-go-build-vs-go-install-c7c0fd135cf9)

- Run directly
  ```
    go run /GatorStore/backend/src/github.com/UF-CEN5035-2022SpringProject/GatorStore/main.go
  ```

- Build into binary file
  ```
    go build github.com/UF-CEN5035-2022SpringProject/GatorStore
  ```
  
  if we face error such as
  ```
    go.mod file not found in current directory or any parent directory; see 'go help modules'
  ```

  try to fix with
  ```
    go env -w GO111MODULE=auto
  ```

- Build module into binary file in bin/
  ```
    go install github.com/UF-CEN5035-2022SpringProject/GatorStore
  ```

# Backend Logger 
1. Import package in the code
For example
```
import (
	"github.com/UF-CEN5035-2022SpringProject/GatorStore/logger"
)
```

2. Three different level Logger to call
- InfoLogger
- WarningLogger
- ErrorLogger

```
logger.InfoLogger.Println(appName + " server is start at port: " + port)
```

3. Output
```
INFO: 2022/02/01 18:16:07 main.go:27: GatorStore server is start at port: 8080
```

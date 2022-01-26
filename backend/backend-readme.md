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

# Run Go program
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

# Variables
---
Variables have been ask to be nice and clean, we are able to add shadow variables, but not redeclare in the same statement.

- Always have to be use
- Declare variables
  ```
  var i int = 2
  ```
  or to let the compiler decide for us, as an auto type in C++
  ```
  i := 2
  ```
- First letter **lower case** will package scope, first letter with upper case to export.
- No private scope.

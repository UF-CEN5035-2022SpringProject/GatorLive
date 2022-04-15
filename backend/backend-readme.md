# Requirements
1. Google API credentials (Key for calling google/youtube api)
   1. Create a new project at https://console.cloud.google.com/
   2. Go to `APIs & Services` section
   3. Enable `YouTube Data API v3`
   4. Create a new `OAuth 2.0 Client ID`
   5. Carefully change the redirect uri as follow and save.(Order matters)
   ![image](https://user-images.githubusercontent.com/11768359/163622161-a9771954-de17-4e68-a3c8-13b6eab26120.png)
   1. Download the credential 
   ![image](https://user-images.githubusercontent.com/11768359/163622397-8764550f-5505-4b02-8d7b-177cbd0242c2.png)

3. Google Database Credentials (Key for remote accessing google firestore)
   1. Create a new project at https://console.firebase.google.com/
   1. Create a new Firestore database
   2. Create all following collections either by code or manually. (`users`,`stores`,`products`,`settings`,`jwtTokenMap`,`lives`)
   3. Under `settings` collection, add following documents(`orderAutoIncrement`,`productAutoIncrement`,`storeAutoIncrement`,`userAutoIncrement`)
   4. In each documents added in previous step, add `number` field.
   5. Go to `Project settings > Service accounts`, generate and download new private key.


5. Golang (v1.17.6)
   - Refer to https://go.dev/doc/install

# Installation

1. Download/Clone the project.
2. Place the `Google API credentials` and `Google Database Credentials` under `GatorStore/backend/src/` directory
3. Opening a command prompt and typing the following command
    ```
    cd GatorStore/backend/src
    go get
    ```

# Run
- Opening a command prompt and typing the following command
    ```
    cd GatorStore/backend/src
    go run main.go
    ```

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
# Dependencies
---
We have create a go.mod in the path below.
```
backend\src\github.com\UF-CEN5035-2022SpringProject\GatorStore
```

Please enter the path and enter ```go get``` to download every dependencies.

# Set Credentials
The credentials have been save in the private Github link below  
https://github.com/UF-CEN5035-2022SpringProject/GatorStoreCredentials

- Google API credentials (Key for calling google api)
- Google Database Credentials (Key for remote accessing google firestore)

If you want to join or test our product please send a email to ***yimingchang@ufl.com*** for develop permission

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

# GatorLive Project Introductions and Rules
---
## Project Members
Canvas Group Link - https://ufl.instructure.com/groups/419331/users   
[Yi-Ming Chang](yimingchang@ufl.edu)  
[Hung-You Chou](hchou@ufl.edu)  
[Vivaan Goomer](vivaangoomer@ufl.edu)  
[Sebastian Llerena](llerenabarruetos@ufl.edu)  

## Introduction
For course CEN5035, our group are building a AmazonLive Clone.
We are using Golang in backend programming and ReactJS in frontend programming.
The purpose is to combine live streaming functions to sell products in users’ online stores. 
We hope to provide users not only a platform to sell their products but also with our features that can help them sell stuff with high efficiency.

Reference: https://www.nngroup.com/articles/livestream-ecommerce-china/
---

## Learnings:
Golang:
Down below are few reference to learn Golang in the first step.

https://go.dev/tour/  
https://go.dev/doc/effective_go  
https://www.youtube.com/watch?v=YS4e4q9oBaU 
 
---

## Github Branching Rules
actions:
- feature
- test
- demo

Backend branch naming
```
spr{#number}-backend-{actions}/{action-name}
```

Seller frontend naming
```
spr{#number}-sfrontend-{actions}/{action-name}
```

Buyer frontend naming
```
spr{#number}-bfrontend-{actions}/{action-name}
```

For example:
spr1-backend-feature/routing

---

## Testing Proccess
1. Self Testing
2. PR into test branch (according to feature into backend-test, seller-frontend-test, buyer-frontend-test)
3. PR into main branch 


---

# Backend
Golang and backend set up please check [backendend-readme.md](https://github.com/UF-CEN5035-2022SpringProject/GatorStore/blob/main/backend/backend-readme.md)

## API Document 
**BACKEND API:** 

 ```http://localhost:8080/{env}/{routePath}```
 - Production {env} = api
   ```
   http://localhost:8080/api/user/login
   ```
 - Test {env} = test/api
   ```
   http://localhost:8080/test/api/user/login
   ```

**Header**  
  | Name | Type | Description |
  | --- | --- | --- |
  | authorization | string | server JWT token, only |
  | time | datetime | string |

**Return Body**
  | Name | Type | Description |
  | --- | --- | --- |
  | status | int | use for passing error code (0 = success, other = error code) |
  | result | json | return object |

- Examples:

  Success 
  ```
  {
      "status": 0,
      "result": {
          ...
      }
  }
  ```
  
  Error
  ```
  {
      "status": 3105, // not 0
      "result": {
          "message": "no store found for user {userId}",
          "type": "",
          "data": null
      }
  }
  ```

### User API URLs
---

#### UA1. User Login API
 - Method: POST
 - {routePath}: /user/login/
 - **Header**
   | Name | Type | Description |
   | --- | --- | --- |
   | time | datetime | request time |
   | Authorization | string | (optional) jwtToken |
   
 - **Request Body Table**
   | Name | Type | Description |
   | ---  | --- | --- |
   | code | string | Oauth2 code for requesting Google API |

    Example:
    1. Without user jwtToken in web page session, login by Oauth2
     ```
     {
         'code': 'qejklsadiup1io135',
     }
     ```

    2. Using user jwtToken login directly, put the token in the header as "Authorization"
     
 - **Response**  
    Success: 
    ```
    {
        "status": 0,
        "result": {
              'id': "113024",
              'name': "YiMing Chang",
              'email': "yimingchang@ufl.edu",
              'jwtToken': "gatorStore_qeqweiop122133"
              'accessToken': "GatorStore_10302323"
              'createTime': "2006-01-02T15:04:05Z07:00"
              'updateTime': "2006-01-02T15:04:05Z07:00"
        }
    }
    ```

    Error:
     ```
     {
         "status": 800,
         "result": {
            "errorName": "MISS_PARAMS"
         }
     }
     ```
   
     Error Code Table for error situation:

      | ErrorName | ErrorCode | HttpStatus | Description |
      | ---  | --- | --- | --- |
      | MISS_PARAMS | 800 | 400 | |
      | INVALID_PARAMS | 801 | 400 | |
      | NO_JWTTOKEN | 1000 | 401 | Missing JwtToken |
      | INVALID_JWTTOKEN | 1001 | 401 | Expire or invalid jwtToken |
      | INVALID_ACCESSTOKEN | 9000 | 403 | Expire Google Access Token |
  
---
#### UA2. Get User Info 
---
#### UA3. User Store List
---
- Method: Get
 - {routePath}: /user/store-list
 - **Header**
   | Name | Type | Description |
   | --- | --- | --- |
   | time | datetime | request time |
   | Authorization | string | jwtToken |
     
 - **Response**  
    Success: 
    ```
    {
        "status": 0,
        "result": {
              "store-list": [
                {
                  'id': "GatorStore_1",
                  'name': "GoGoGator",
                  'userId': "11001",
                  'createTime': "2006-01-02T15:04:05Z07:00",
                  'updateTime': "2006-01-02T15:04:05Z07:00",
                  'isLive': True
                },
                {
                  'id': "GatorStore_1",
                  'name': "IamTheHero",
                  'userId': "11001",
                  'createTime': "2006-01-02T15:04:05Z07:00",
                  'updateTime': "2006-01-02T15:04:05Z07:00",
                  'isLive': True
                },
              ]
        }
    }
    ```

    Error:
     ```
     {
         "status": 800,
         "result": {
            "errorName": "MISS_PARAMS"
         }
     }
     ```
   
     Error Code Table for error situation:

      | ErrorName | ErrorCode | HttpStatus | Description |
      | ---  | --- | --- | --- |
      | MISS_PARAMS | 800 | 400 | |
      | INVALID_PARAMS | 801 | 400 | |
      | NO_JWTTOKEN | 1000 | 401 | empty jwtToken |
      | INVALID_JWTTOKEN | 1001 | 401 | Expire or invalid jwtToken |
      | INVALID_ACCESSTOKEN | 9000 | 403 | Expire Google Access Token |


### Store API URLs
---
#### SA1. Store Livestream API
 - Method: POST
 - {routePath}: /store/{store id}/livestream
 - **Header**
   | Name | Type | Description |
   | --- | --- | --- |
   | Authorization | string | Use for GatorStore Login |
   
 - **Request Body Table**
   | Name | Type | Description |
   | ---  | --- | --- |
   | title | string | Use for naming the new livestream |
   
    Example:
    1. Using user jwtToken login directly
     ```
     {
         'title': '123'
     }
     ```
     
 - **Response**  
    Success: 
    ```
    {
        "status": 0,
        "result": {
              'id': "113024", // livestream id
              'title': "YiMing Chang", // livestream title
              'streamKey': "1324-5678-8974-1230",
              'streamUrl': "some url",
              'createTime': "2006-01-02T15:04:05Z07:00"
              'updateTime': "2006-01-02T15:04:05Z07:00"
              'embedHTML': "some iframe html"
        }
    }
    ```

    Error:
     ```
     {
         "status": 800,
         "result": {
            "errorName": "MISS_PARAMS"
         }
     }
     ```
   
     Error Code Table for error situation:

      | ErrorName | ErrorCode | HttpStatus | Description |
      | ---  | --- | --- | --- |
      | MISS_PARAMS | 800 | 400 | |
      | INVALID_PARAMS | 801 | 400 | |
      | NO_JWTTOKEN | 1000 | 400 | |
      | INVALID_JWTTOKEN | 1001 | 401 | Expire or invalid jwtToken |
      | INVALID_ACCESSTOKEN | 9000 | 403 | Expire Google Access Token |
      
#### SA2. Store Livestream status API
 - Method: GET/PUT
 - {routePath}: /store/{store id}/livestreamStatus
 - **Header**
   | Name | Type | Description |
   | --- | --- | --- |
   | Authorization | string | Use for GatorStore Login |
   
 - **Request Body Table**
   | Name | Type | Description |
   | ---  | --- | --- |
   | status | boolean | streamStatus |
   
    PUT Example:
     ```
     {
         'isLive': true
     }
     ```
    GET Example:
    ```
    {
    }
    ```
 - **Response**  
    Success: 
    ```
    {
      'id': "GatorStore_1",
      'name': "GoGoGator",
      'userId': "11001",
      'createTime': "2006-01-02T15:04:05Z07:00",
      'updateTime': "2006-01-02T15:04:05Z07:00",
      'isLive': True // might change
    }
    ```

    Error:
     ```
     {
         "status": 800,
         "result": {
            "errorName": "MISS_PARAMS"
         }
     }
     ```
   
     Error Code Table for error situation:

      | ErrorName | ErrorCode | HttpStatus | Description |
      | ---  | --- | --- | --- |
      | MISS_PARAMS | 800 | 400 | |
      | INVALID_PARAMS | 801 | 400 | |
      | NO_JWTTOKEN | 1000 | 400 | |
      | INVALID_JWTTOKEN | 1001 | 401 | Expire or invalid jwtToken |
      | INVALID_ACCESSTOKEN | 9000 | 403 | Expire Google Access Token |
  
### Product API URLs
---

### Test API URLs
---
## Ojbect Document
---
Object Table Columns 

| Name | Description |
| ---  | --- | 
| Var/Method | Showing this attibute is a variable or method |
| Key/Optional | - 'K' represents primary key.<br /> - 'RK means it's related with others to combine into primary key.<br /> - 'O' means optional |
| Type | variable type or method return type |
| Description | descibe the attribute purpose |

---
#### JWT Object
This object will be use for respresenting the user in server, every user has an unique JWT token for requesting API.
Use it as Authentication in the middleware of the server (before the request passby).
If the token object does not match to the request, return httpStatus 401 or 403.

The key will be ***jwtToken***, and the value will be as below:
| Var/Method | Key/Optional | Type | Description |
| ---  | --- | --- | --- |
| jwtToken | k | string | unique jwt authorization key in GatorStore|
| email | | string | unique identifier |
| createTime |  | string | create datetime | 

JSON Example:
```
{
  'email': "yimingchang@ufl.edu",
  'jwtToken': "gst.R2F0b3JTdG9yZV95aW1pbmdzdGFyNTU2NkBnbWFpbC5jb20xMTAwMg==_MjAyMi0wMi0yMlQwMjoyNTowMVo=",
  'createTime': "2006-01-02T15:04:05Z07:00"
}
```

---
#### User Object
If user has already register, an **uniqueId** will be assigned to user.  
Or else we'll use **email** as a identifier.

The key will be **email**, and the value will be as below:
| Var/Method | Key/Optional | Type | Description |
| ---  | --- | --- | --- |
| id | K | string | userId - unique identifier |
| name |   | string | Receive by google api |
| email | K | string | unique identifier |
| jwtToken |  | string | unique jwt authorization key in GatorStore |
| accessToken |  | string | youtube access token | 
| createTime |  | string | create datetime | 
| updateTime |  | string | latest update datetime | 

JSON Example:
```
{
  'id': "113024",
  'name': "YiMing Chang",
  'email': "yimingchang@ufl.edu",
  'jwtToken': "gst.R2F0b3JTdG9yZV95aW1pbmdzdGFyNTU2NkBnbWFpbC5jb20xMTAwMg==_MjAyMi0wMi0yMlQwMjoyNTowMVo=",
  'accessToken': "GatorStore_10302323",
  'createTime': "2006-01-02T15:04:05Z07:00"
  'updateTime': "2006-01-02T15:04:05Z07:00"
}
```
---
#### Store Object
Each Store will have uniqueId, and belong to one user who created it.

The key will be **storeId**, and the value will be as below:
| Var/Method | Key/Optional | Type | Description |
| ---  | --- | --- | --- |
| id | K | string | storeId - unique identifier |
| name |   | string | storeName |
| userId | | string | unique creator |
| createTime |  | string | create datetime | 
| updateTime |  | string | latest update datetime | 
| isLive | | boolean | check if this store is on live |
| liveId | | string | liveObj's ID to get live information, empty if isLive is false |

JSON Example:
```
{
  'id': "GatorStore_1",
  'name': "GoGoGator",
  'userId': "11001",
  'createTime': "2006-01-02T15:04:05Z07:00",
  'updateTime': "2006-01-02T15:04:05Z07:00",
  'isLive': True
  'liveId': "132001"
}
```
---

#### Live Object 
The key will be **liveId**, and the value will be as below:
| Var/Method | Key/Optional | Type | Description |
| ---  | --- | --- | --- |
| id | K | string | live - unique identifier for liveObj |
| storeId | K | string | store creating this live |
| title |   | string | storeName |
| streamKey |  | string | create datetime | 
| streamUrl |  | string | latest update datetime | 
| createTime | | datetime | live create time |
| updateTime | | datetime | live update time |
| embedHTML | | string | use for web iframe ebmed |

 ```
 {
   'id': "113024", // livestream id
   'title': "YiMing Chang", // livestream title
   'storeId': "122323"
   'streamKey': "1324-5678-8974-1230",
   'streamUrl': "some url",
   'createTime': "2006-01-02T15:04:05Z07:00"
   'updateTime': "2006-01-02T15:04:05Z07:00"
   'embedHTML': "some iframe html"
 }
 ```
 
---
### Global ErrorCode  
- General Errors such as missing params or invalid params will be under 1000
- Google Errors error code will be starting with 9xxx
- GatorStore Errors will be in range 1000 ~ 8xxx

- Server Request Errors
   | ErrorName | ErrorCode | Description |
   | ---  | --- | --- |
   | UnknownInternalErrCode | 800 |  |
   | MissingParamsCode | 801 |  | 
   | InvalidParamsCode | 802 |  | 
   
- DB errors
   | ErrorName | ErrorCode | Description |
   | ---  | --- | --- |
   | UnknownDbErrCode | 900 | |
   | UnableToGetDbObj | 901 | |
 
- JWT errors
   | ErrorName | ErrorCode | Description |
   | ---  | --- | --- |
   | MissingJwtTokenCode | 1000 | | Empty jwtToken |
   | InvalidAccessTokenCode | 1001 |  | Expire or invalid jwtToken |

- Error with Google
   | ErrorName | ErrorCode | Description |
   | ---  | --- | --- |
   | MissingAccessTokenCode | 9000 | |
   | InvalidAccessTokenCode | 9001 | |
   | InvalidGoogleCode | 9002 | |
   
 

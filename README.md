# GatorLive Project Introductions and Rules
---
## Project Members
Canvas Group Link - https://ufl.instructure.com/groups/419331/users   
[Yi-Ming Chang](mailto:yimingchang@ufl.edu)  
[Hung-You Chou](mailto:hchou@ufl.edu)  
[Sebastian Llerena](mailto:llerenabarruetos@ufl.edu)  

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
              'accessToken': "GatorStore_10302323" (remove?)
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
#### UA2 User Info API

Get user info
- Method: GET
 - {routePath}: /user/{userId}/info

 - **Header**
   | Name | Type | Description |
   | --- | --- | --- |
   | Authorization | string | Use for GatorStore Login |
 
 - **Response**  
    Success: 
    ```
    {
      "status": 0,
      "result": {
          "accessToken": "dfsdfsdf",
          "createTime": "2022-02-22T02:25:01Z",
          "email": "yimingstar5566@gmail.com",
          "id": "11002",
          "jwtToken": "gst.R2F0b3JTdG9yZV95aW1pbmdzd",
          "name": "Yi-Ming Chang",
          "updateTime": "2022-02-22T02:25:01Z"
      }
    }
    ```

    Error:
     ```
     {
         "status": 801,
         "result": {
            "errorName": "MissingParamsCode"
         }
     }
     ```
   
     Error Code Table for error situation:
     
     | ErrorName | ErrorCode | HttpStatus | Description |
     | ---  | --- | --- | --- |
     | UnknownInternalErrCode | 800 | 500 | |
     | MissingParamsCode | 801 | 400 | |
     | InvalidParamsCode | 802 | 403 | |
     | MissingJwtTokenCode | 1000 | 401 | |
     | InvalidJwtTokenCode | 1001 | 401 | Expire or invalid jwtToken |

---
#### UA3 User Store List API

Get the users store list, split the item with page
 - Method: GET
 - {routePath}: /user/{userId}/store-list?page={page}
   - page parameter decide which page requesting, if overflow, return the last page. If missing, return page 0.
 
 - **Header**
   | Name | Type | Description |
   | --- | --- | --- |
   | Authorization | string | Use for GatorStore Login |
   
- **Request Body Table**   
    Empty request body
    GET Example:
    ```
    {
    }
    ```
- **Response** 
    If storeList is empty, the value will be null 
    Success: 
    ```
    {
      "status": 0,
      "result": {
          "currectPage": 0,
          "maxPage": 0,
          "storeList": [
              {
                  "createTime": "2022-03-29T02:32:31Z",
                  "id": "gatorstore-1",
                  "isLive": false,
                  "liveId": "",
                  "name": "YiMing Fantastic Store",
                  "updateTime": "2022-03-29T02:32:31Z",
                  "userId": "11002"
              }, ...
          "userId": "11002"
      }
    }
    ```

    Error:
     ```
     {
         "status": 1000,
         "result": {
            "errorName": "MissingJwtTokenCode"
         }
     }
     ```
   
     Error Code Table for error situation:

      | ErrorName | ErrorCode | HttpStatus | Description |
      | ---  | --- | --- | --- |
      | MissingJwtTokenCode | 1000 | 401 | |
      | InvalidJwtTokenCode | 1001 | 401 | Expire or invalid jwtToken |
     
### Store API URLs
---
#### SA0. Store Create API

 - Method: POST
 - {routePath}: /store/create
 - **Header**
   | Name | Type | Description |
   | --- | --- | --- |
   | Authorization | string | Use for GatorStore Login |

 - **Request Body Table**
   | Name | Type | Description |
   | ---  | --- | --- |
   | name | string | unique name for store, cannot be empty|
   
   Example:
   ```
     {
         'name': '123'
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
     'isLive': False // might change
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

---
#### SA1. Get Store Object API
 - Method: GET
 - {routePath}: /store/{storeId}/info

 - **Header** - Empty
 
 - **Response**  
    Success: 
    ```
    {
      'id': "GatorStore_1",
      'name': "GoGoGator",
      'userId': "11001",
      'createTime': "2006-01-02T15:04:05Z07:00",
      'updateTime': "2006-01-02T15:04:05Z07:00",
      'isLive': True
      'liveId': "live_133312"
    }
    ```

    Error:
     ```
     {
         "status": 801,
         "result": {
            "errorName": "MissingParamsCode"
         }
     }
     ```
   
     Error Code Table for error situation:
     
     | ErrorName | ErrorCode | HttpStatus | Description |
     | ---  | --- | --- | --- |
     | UnknownInternalErrCode | 800 | 500 | |
     | MissingParamsCode | 801 | 400 | |
     | InvalidParamsCode | 802 | 403 | |
     | MissingJwtTokenCode | 1000 | 401 | |
     | InvalidJwtTokenCode | 1001 | 401 | Expire or invalid jwtToken |
   
---       
#### SA2. Store Product List API
Get the products according to the store, split the item with page
 - Method: GET
 - {routePath}: /store/{storeId}/product-list?page={page}
   - page parameter decide which page requesting, if overflow, return the last page. If missing, return page 0.
 
 - **Header**  - Empty
   
- **Request Body Table**   
    Empty request body
    GET Example:
    ```
    {
    }
    ```
- **Response**  
    If productList is empty, the value will be null
    Success: 
    ```
    {
      'storeId': "GatorStore_1",
      'maxPage': 3, (start at 0)
      'currentPage': 0,
      'productList': [
        {productObject},
        {productObject},
        ...
      ]
    }
    ```

    Error:
     ```
     {
         "status": 1000,
         "result": {
            "errorName": "MissingJwtTokenCode"
         }
     }
     ```
   
     Error Code Table for error situation:

      | ErrorName | ErrorCode | HttpStatus | Description |
      | ---  | --- | --- | --- |
      | MissingJwtTokenCode | 1000 | 401 | |
      | InvalidJwtTokenCode | 1001 | 401 | Expire or invalid jwtToken |
      
---
#### SA3. Store OrderList API
Get the orders according to the store, split the item with page
 - Method: GET
 - {routePath}: /store/{storeId}/order-list?page={page}
   - page parameter decide which page requesting, if overflow, return the last page. If missing, return page 0.
 
 - **Header**
   | Name | Type | Description |
   | --- | --- | --- |
   | Authorization | string | Use for GatorStore Login |
   
- **Request Body Table**   
    Empty request body
    GET Example:
    ```
    {
    }
    ```
- **Response**  
    Success: 
    ```
    {
      'storeId': "GatorStore_1",
      'maxPage': 3, (start at 0)
      'currentPage': 0,
      'orderList': [
        {orderObject},
        {orderObject},
        ...
      ]
    }
    ```

    Error:
     ```
     {
         "status": 1000,
         "result": {
            "errorName": "MissingJwtTokenCode"
         }
     }
     ```
   
     Error Code Table for error situation:
      | ErrorName | ErrorCode | HttpStatus | Description |
      | ---  | --- | --- | --- |
      | MissingJwtTokenCode | 1000 | 401 | |
      | InvalidJwtTokenCode | 1001 | 401 | Expire or invalid jwtToken |

--- 
### Live API URLs
---
#### SLA0. Store Livestream API
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
      
---
#### SLA1. Store Livestream status API
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
      
---
### Product API URLs
---
#### PA1. Product Create API
 - Method: POST
 - {routePath}: /product/create
 - **Header**
   | Name | Type | Description |
   | --- | --- | --- |
   | Authorization | string |  jwtToken |
   
 - **Request Body Table**
   | Name | Type | Description |
   | ---  | --- | --- |
   | name | string | product name |
   | price | float64 | product price |
   | description | string | product introduction |
   | quantity | int | how many product in stock |
   | picture | string | picture in base64 |
   | storeId | string | publish to which store |

    Example:
     ```
     {
         'name': 'gator',
         'price": 1000,
         'description':'real gator',
         'quantity': 1,
         'picture':'123123123',
         'storeId':'11001'
     }
     
 - **Response**  
    Success: 
    ```
    {
        "status": 0,
        "result": {
              'id': "113024", // product id
              'name':'gator',
              'storeId':'11001',
              'createTime': "2006-01-02T15:04:05Z07:00",
              'updateTime': "2006-01-02T15:04:05Z07:00",
              'price':'1000',
              'quantity':'1',
              'description':'real gator',
              'picture':'123123',
              'isDeleted':false
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
#### PA2. Product Get API
 - Method: GET
 - {routePath}: /product/{product id}/info
 - **Header** - Empty
 - **Request Body Table**
   | Name | Type | Description |
   | ---  | --- | --- |

    Example:
     ```
     {
     }
     
 - **Response**  
    Success: 
    ```
    {
        "status": 0,
        "result": {
             {productObject}
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
#### PA3. Product Update API
 - Method: PUT
 - {routePath}: /product/{product id}
 - **Header**
   | Name | Type | Description |
   | --- | --- | --- |
   | Authorization | string |  jwtToken |
   
 - **Request Body Table**
   | Name | Type | Description |
   | ---  | --- | --- |
   | (optional)name | string | product name |
   | (optional)price | string | product price |
   | (optional)description | string | product introduction |
   |(optional) quantity | string | how many product in stock |
   | (optional)picture | string | picture in base64 |

    Example:
     ```
     {
         'name': 'gator',
         'price": '1000',
         'description':'real gator',
         'quantity': '1',
         'picture':'123123123'
     }
     
 - **Response**  
    Success: 
    ```
    {
        "status": 0,
        "result": {
              {productObject}
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
#### PA4. Product Purchase API
 - Method: POST
 - {routePath}: /product/{product id}
 - **Header**
   | Name | Type | Description |
   | --- | --- | --- |
   | Authorization | string |  jwtToken |
   
 - **Request Body Table**
   | Name | Type | Description |
   | ---  | --- | --- |
   |quantity | string | how many product in stock |
   |liveId| string | bind with livestream |

    Example:
     ```
     {
         'quantity': '1'
         'liveId': ''  (optional)
     }
     
 - **Response**  
    Success: 
    ```
    {
      "status": 0,
      "result": {
          "createTime": "2022-03-30T14:58:43Z",
          "id": "order-5",
          "liveId": "",
          "productId": "product-1",
          "quantity": 1,
          "storeId": "gatorstore-1",
          "subtotal": 1000,
          "userId": "11002"
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
#### PA5. Product Delete API
 - Method: DELETE
 - {routePath}: /product/{product id}
 - **Header**
   | Name | Type | Description |
   | --- | --- | --- |
   | Authorization | string |  jwtToken |
   
 - **Request Body Table**
   | Name | Type | Description |
   | ---  | --- | --- |

    Example:
     ```
     {
     }
     
 - **Response**  
    Success: 
    ```
    {
        "status": 0,
        "result": {
              'id': "113024", // product id
              'name':'gator',
              'storeId':'11001',
              'createTime': "2006-01-02T15:04:05Z07:00",
              'updateTime': "2006-01-02T15:04:05Z07:00",
              'price':'1000',
              'quantity':'1',
              'description':'real gator',
              'picture':'123123',
              'isDeleted':true
        }
    }
    ```

    Error:
     ```
     {
         "status": 800,
         "result": {
            "errorName": "INVALID_JWTTOKEN"
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
| liveId | | string | liveObj's ID to get live information, empty string if isLive is false |

JSON Example:
```
{
  'id': "GatorStore_1",
  'name': "GoGoGator",
  'userId': "11001",
  'createTime': "2006-01-02T15:04:05Z07:00",
  'updateTime': "2006-01-02T15:04:05Z07:00",
  'isLive': True
  'liveId': "" 
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
| embedChatRoom | | string | use for web iframe ebmed |

 ```
 {
   'id': "113024", // livestream id
   'title': "YiMing Chang", // livestream title
   'storeId': "122323"
   'streamKey': "1324-5678-8974-1230",
   'streamUrl': "some url",
   'createTime': "2006-01-02T15:04:05Z07:00",
   'updateTime': "2006-01-02T15:04:05Z07:00",
   'embedHTML': "some iframe html",
   'embedChatRoom': "chatroom iframe html"
 }
 ```
---
#### Product Object
   | Name | Type | Description |
   | ---  | --- | --- |
   | id | string | product id |
   | name | string | product name |
   | price | float| product price |
   | description | string | product introduction |
   | quantity | int| how many product in stock |
   | picture | string | picture in base64 |
   | storeId | string | publish to which store |
   | createTime | string | createTime |
   | updateTime | string | updateTime|
   | isDeleted | boolean | |
   ```
   {
       'id': "113024", // product id
       'name':'gator',
       'storeId':'11001',
       'createTime': "2006-01-02T15:04:05Z07:00",
       'updateTime': "2006-01-02T15:04:05Z07:00",
       'price':'1000',
       'quantity':'1',
       'description':'real gator',
       'picture':'123123',
       'isDeleted':false
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
   | InvalidJwtTokenCode | 1001 |  | Expire or invalid jwtToken |

- Error with Google
   | ErrorName | ErrorCode | Description |
   | ---  | --- | --- |
   | MissingAccessTokenCode | 9000 | |
   | InvalidAccessTokenCode | 9001 | |
   | InvalidGoogleCode | 9002 | |

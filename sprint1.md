# Sprint1
---
## Demo Video
https://youtu.be/UeHNUq7gqqM

## User Stories
---
Our features are according to user stories.
We use [Github Project Board](https://github.com/orgs/UF-CEN5035-2022SpringProject/projects/1) to record it down.

- **(Sprint1) Seller System**  
  In the platform, We hope there's a friendly user interface that can provide us with an easy way to upload our product information. 
  We can directly use the platform to manage our commodities.
  Additionally, We hope we could also show our offline store information. In this way, this online store will greatly benefit our brand. 
  
## Features
---
- User System (Sprint1)
  - Sellers must log in (Gmail Login)
  - Unique ID matching sellers and sellers are able to create stores

- E-commerce System (Sprint1)
  - Sellers can create products in E-commerce
     - Amount
     - Description
     - Image

## Issues
---
We use project board to define issue into 3 stage
- ```To do```: after discussion we create issue card and assign to members.
- ```In Progress```
- ```Done```: When the issue is close we'll move it to Done.  

![ProjectBoardIssue](https://user-images.githubusercontent.com/69064626/152628222-d657b42a-ef3b-4bd1-9872-9bd0ef424b5a.png)

Accordingly, each component or task will be create into issue and it's also important for us to announce bug as a issue during the development.
We use different labels to recognize the responsibility.  

![IssueList](https://user-images.githubusercontent.com/69064626/152628355-15e98d74-6767-4676-b524-ef90a468e24a.png)

All the discussions and even surveys we record on the issue to acknowledge all the members.
For example, 
- [Google Login Survey](https://github.com/UF-CEN5035-2022SpringProject/GatorStore/issues/7)
- [User API Design & User Object design](https://github.com/UF-CEN5035-2022SpringProject/GatorStore/issues/12)

## [API Documentations](https://github.com/UF-CEN5035-2022SpringProject/GatorStore#api-document)
---
The most critical part is to make a explicit document for the API integration between backend and frontend.
In the section, we define the request routing url, request body, request method and response type. 

For example, our google login api ([UA1](https://github.com/UF-CEN5035-2022SpringProject/GatorStore#ua1-user-login-api)) is define below as the image shows.

<img width="897" alt="UA1" src="https://user-images.githubusercontent.com/69064626/152628543-2928e86d-272f-4362-a7cf-f0a291fb3a5c.png">

## Structure of the Project
---
### Backend (Golang)
  - [gorilla/mux](https://github.com/gorilla/mux) - use for routing and easy to add middleware
    - [Survey on routing package](https://github.com/UF-CEN5035-2022SpringProject/GatorStore/issues/19)   
  - FireStore - NoSQL provided by google. With 20000 read/write quata per day for free.
    - [Survey on database](https://github.com/UF-CEN5035-2022SpringProject/GatorStore/issues/4)  
  
### Frontend
  - MaterialUI
  - ReactJS
  - ReackHooks

## Apendix - Document on Setting Develop Environment
---
- [Backend Document](https://github.com/UF-CEN5035-2022SpringProject/GatorStore/blob/main/backend/backend-readme.md)
  - Logger
  - DB settings
  - Package Dependencies   

- Tools
  - Postman: able to make APIs into collection.  
    ![PostMan](https://user-images.githubusercontent.com/69064626/152628849-72a8efd3-fdcc-4861-b1ef-5dfbe2e19935.png)

# Sprint3
---
## What we have done in Sprint3
- More New Features
  - https://github.com/UF-CEN5035-2022SpringProject/GatorStore/blob/main/sprint3.md#new-feautures
- Our Project Board
  - https://github.com/UF-CEN5035-2022SpringProject/GatorStore/blob/main/sprint3.md#project-board
- Instructions on running the applications are in each applications
  - GatorStore Backend  
    https://github.com/UF-CEN5035-2022SpringProject/GatorStore/blob/main/backend/backend-readme.md
  - Seller E-Commerse Manage Web application  
    https://github.com/UF-CEN5035-2022SpringProject/GatorStore/blob/main/seller-frontend/seller-frontend-readme.md
  - Buyer E-Commerse Web application
    https://github.com/UF-CEN5035-2022SpringProject/GatorStore/blob/main/buyer-frontend/buyer-frontend-readme.md
- Demo Video
  - https://github.com/UF-CEN5035-2022SpringProject/GatorStore/blob/main/sprint3.md#demo-video
- More Test Case for new features
  - https://github.com/UF-CEN5035-2022SpringProject/GatorStore/blob/main/sprint3.md#testing

---
## Demo Video
https://www.youtube.com/watch?v=ioquXggbBkY

---
## Project Board
https://github.com/orgs/UF-CEN5035-2022SpringProject/projects/1

Down below link is the cards with sprint3 label  
https://github.com/orgs/UF-CEN5035-2022SpringProject/projects/1?card_filter_query=label%3Asprint3

---
## New Feautures

### User Stories
Just for remind, We have 2 web applications   
  - E-commerse manage system for Seller. 
  - E-commerse web application for Buyers.  
    
Our purpose is to combine live stream platform and let Seller to sell their products by live stream.  
Eventually, the Buyer (audience) can buy products from the E-commerse because of watching the live.  
Or simply, just buy things when they are browsing the stores.  

So, most of the features in this sprint we focus on the store, live and orders for user.  

### 1. Backend APIs & Frontend
  - Implement RESTful apis for Store, Product.
    - Check out issue #76
    - Backend api integration with frontend
    - Create order history after purchasing a product
  - Link id of new youtube livestream with the store 
  - Add promotion product list to the livestream. issue #83
  - Remove authentication of a few APIs, issue #76
  - Implement Product page
  - Implement Buyer webpage
  - Expend the verification of jwtToken and users' accessibility of backend apis. #100 #88
      - Passing user object in the middleware down with the request.
    
### 2. Frontend Components     
(the API's implemented in the front-end can be found on issue #76 and #83)
  - Create a new React Application for the buyer, with some of the following features
  - Get product list for a store via pages of 5 elements each
    - For this and any other paging retrievals, we make the page load the next page of 5 items when the user scrolls down to the bottom.
    - ![image](https://user-images.githubusercontent.com/40399062/161351035-75863e79-94ef-42d4-aa0e-dbd3c5400952.png)
  - Get store list for a seller account via pages of 5 elements each
  - Get product information for a product via its URL
    - URL "product/product-id" will automatically retrieve the data for product-id using a GET API
    - ![image](https://user-images.githubusercontent.com/40399062/161351157-07040c86-7768-46bb-8e87-9268627077ab.png)
  - Get store information for a product via URL
    - URL "store/store-id" will automatically get the data for a store, which also includes is live status and embedded livestream (if live), using a GET API
  - For a store, the livestream API is called every 5 seconds to get current "isLive" status
    - If this changes, the page will automatically adapt its live component to show the embedded livestream (both video and chat)
  - Seller can select "Featured Items" when creating a livestream. These get sent to backend and get retrieved for the viewers
    - Viewers can click on these items to go to their pages and purchase them
  - Allow seller to create new products
  - Implemented the following actions for a seller on a product page:
      - Update Product data
      - Delete Product
  - Purchasing a product (buyer action) from a product page of 5 elements each
  - Displaying the list of orders for a buyer using pages of 5 elements each
  - On product page, linked button to go back to its store (store/store-id)
  - On store page, linked buttons to go to a product's page (product/product-id)
  - Added a list of recommended stores on landing page
    - An API populates this list with every store available. Will implement a more streamlined and system for this next sprint.
  - Embed youtube live chat iframe next to youtube livestream video iframe. 
      - Two options while we are designing - Check this #74
          1. Build our own chat room, implement product purchase by user comments
          2. Use Youtube live chat iframe, and build a product list which can be shown specifically during live - ***Feature Items***
            <br></br>
            <img src="https://user-images.githubusercontent.com/69064626/161345098-ad18bad8-ab4e-420e-90cd-9fcda8ad152b.jpeg" width="800">     
            
            
(The list and progress of the APIs integrated were recorded in issue #93)


### 3. Backend Testing
  - Check this section https://github.com/UF-CEN5035-2022SpringProject/GatorStore/blob/main/sprint3.md#testing
---
## Bug Fix
1. APIS & Frontend
    - Youtube livestream iframe showing errors. issue #72
    - Youtube livestream not starting even when livestream signal pushed to youtube. issue #98
    - Backend api return multiple response. issue #102
## Testing
1. Golang Backend testing (Because there are conflictions between libraries, we are not able to test thoroughly.)
    - Focus on the http tesing by using [httptest](https://pkg.go.dev/net/http/httptest)
    - Fixed testcases for modified apis(get livestream status)
    - Added total of 14 testcases for product and store apis
        - Create Product
        - Get Product info
        - Create Store
        - Get Store info
        - Get Product list of a Store
 
## Future Work in Sprint 4
- Dockerize Application
- Auntomatic Test by using PostMan API (For example, Jenkins can import Postman file and run all API by real request)
- More Instruction for general user
- Design explicit instructions on Landing Page --> More comprehensive prototype

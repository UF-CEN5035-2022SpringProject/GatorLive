# Sprint3
---
## Demo Video

---
## Project Board
https://github.com/orgs/UF-CEN5035-2022SpringProject/projects/1

Down below link is the cards with sprint3 label  
https://github.com/orgs/UF-CEN5035-2022SpringProject/projects/1?card_filter_query=label%3Asprint3

---
## Application Instruction Document

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
  - Embed youtube live chat iframe next to youtube livestream video iframe. 
      - Two options while we are designing - Check this #74
          1. Build our own chat room, implement product purchase by user comments
          2. Use Youtube live chat iframe, and build a product list which can be shown specifically during live - ***Feature Items***
            <br></br>
            <img src="https://user-images.githubusercontent.com/69064626/161345098-ad18bad8-ab4e-420e-90cd-9fcda8ad152b.jpeg" width="800">



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
 

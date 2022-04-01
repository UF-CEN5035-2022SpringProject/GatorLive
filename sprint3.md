# Sprint3
---
## Demo Video

---
## Application Instruction Document

## New Feautures
1. APIs & Frontend
    - Implement RESTful apis for Store, Product.
      - Check out issue #76
      - Backend api integration with frontend
      - Create order history after purchasing a product
    - Link id of new youtube livestream with the store 
    - Add promotion product list to the livestream
    - Remove authentication of a few APIs, issue #76
    - Implement Product page
    - Implement Buyer webpage
    - Add verification of jwtToken and users' accessibility of backend apis
    
2. Embed iframe
    - Embed youtube live chat next to youtube livestream video iframe.
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
 

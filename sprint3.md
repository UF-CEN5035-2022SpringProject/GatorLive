# Sprint3
---
## Demo Video

---
## Application Instruction Document

## New Feautures
1. APIs & Frontend
    - Implement RESTful apis for Store, Product and Order objects.
      - Check out issue [#76](https://github.com/UF-CEN5035-2022SpringProject/GatorStore/issues/76)
      - Backend api integration with frontend
    - Implement Product page
    
2. Embed iframe
    - Embed youtube live chat next to youtube livestream video iframe.

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
 

# Sprint 2
## Demo Video

## Project Board
 
---
## Features

1. APIs
   - Integrate Youtube APIs
   - Integrate Live APIs
      - Start & End Live
      - Live Status   
   - JWT token in APIs header
2. Authentications
3. Embed Iframe in Frontend (Requires timing, but this issue is being explored in issue [#72](https://github.com/UF-CEN5035-2022SpringProject/GatorStore/issues/72))

4. Build the API document format
---
## Testing
1. Golang Backend Testing
2. Cypress Frontend Testing (using "cypress-social-logins" library with problems; these are being explored in issue [#67](https://github.com/UF-CEN5035-2022SpringProject/GatorStore/issues/67))


---
## Future Work in Sprint3
1. Design Buyer Frontend Page
   - Product List Page 
      - Watch the "live" created by the store and able to buy the products on the list
   - Create Order System
   - Finish the Product System
2. Solve the automation testing in frontend 
   - Unable to use Cypress to test Google Login, blocked by Google server
 
3. Fix the embeded iframe issue (We have to wait 15s till the Youtube server get the stream)
   - Might want to keep checking the youtube channel status and show the iframe in proper time

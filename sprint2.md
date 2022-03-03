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
3. Embed Iframe in Frontend
   - Logged in seller clicks on "Start Livestream" and selects a title and products to advertise
   - Back-end receives user information and title, which it uses to create livestream
   - YouTube API returns the stream's key, iframe embed code, and URL, which the back-end sends to the front-end to be displayed for the user
   - User inputs URL and key into OBS Studio and starts livestreaming
   - After some seconds, user clicks on "GO LIVE" to see the iframe embedded
   - User can click on "End Livestream" to return to "not-live" state
For a more detailed explanation, refer to issue [that specifies the flow and features of Sprint 2](https://github.com/UF-CEN5035-2022SpringProject/GatorStore/issues/54)

4. Build the API document format
---
## Testing
1. Golang Backend Testing
   Check out more details in this issue -
   [Build backend server test](https://github.com/UF-CEN5035-2022SpringProject/GatorStore/issues/27)

   We'll do these tests in the order below: 
   - Unit test on database operations
   - Unit test on Login and LiveStream APIs
     Since the database operations use by these APIs have been tested. We'll focus on the http tesing by using [httptest](https://pkg.go.dev/net/http/httptest)

   
2. Cypress Frontend Testing (using "cypress-social-logins" library with problems; these are being explored in issue [#67, which details how Cypress is unable to get past the Google authentification, as well as the approaches attempted](https://github.com/UF-CEN5035-2022SpringProject/GatorStore/issues/67))


---
## Future Work in Sprint3
1. Design Buyer Frontend Page
   - Product List Page 
      - Watch the "live" created by the store and able to buy the products on the list
   - Create Order System
   - Finish the Product System
2. Solve the automation testing in frontend 
   - Unable to use Cypress to test Google Login, blocked by Google server 
     - The next approach that will be taken will be using Selenium. This is explained in issue [#72](https://github.com/UF-CEN5035-2022SpringProject/GatorStore/issues/72)
 
3. Fix the embeded iframe issue (We have to wait 15s till the Youtube server get the stream)
   - Might want to keep checking the youtube channel status and show the iframe in proper time

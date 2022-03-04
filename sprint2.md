# Sprint 2
## Demo Video
https://www.youtube.com/watch?v=9eUHQexsQaA

## Project Board
https://github.com/orgs/UF-CEN5035-2022SpringProject/projects/1

Down below link is the cards with sprint2 label  
https://github.com/orgs/UF-CEN5035-2022SpringProject/projects/1?card_filter_query=label%3Asprint2

## What we have done in Sprint2
- Backend and frontend are integrated
  - Using RESTful API with JSON format to communicate
- Tesing for both frontend and backend - Check [Testing Chapter](https://github.com/UF-CEN5035-2022SpringProject/GatorStore/blob/main/sprint2.md#testing) for more details
  - Cypress + unit tests in the framework you are using
    - Automation testing on Google Login and Live Streaming page
    - Unit test the button and components on 2 major pages
  - Unit tests for the backend
    - Using go test to test dp operations and Login and Live APIs

- Documentation for the backend API developed - [API document](https://github.com/UF-CEN5035-2022SpringProject/GatorStore#api-document)

---
## New Features
1. APIs & Frontend
   - Integrate Youtube Live Steam APIs and Implement Youtube Iframe in Frontend 
     - Check out issue [#54](https://github.com/UF-CEN5035-2022SpringProject/GatorStore/issues/54) 
     - Backend api integration with frontend and also youtube APIs
   - Integrate Live APIs
     - Check out issue [#45](https://github.com/UF-CEN5035-2022SpringProject/GatorStore/issues/45)
     - Start & End Live
     - Live Status   
   - JWT token in APIs header for **Authentications** 
     - Adding JWT object to database and require in frontend request header
2. User Authentications
   - Build JWT authentication in the middleware - check issue [Build Server JWT generator and authorization functions](https://github.com/UF-CEN5035-2022SpringProject/GatorStore/issues/51)

3. Embed Iframe in Frontend
   - Logged in seller clicks on "Start Livestream" and selects a title and products to advertise
   - Back-end receives user information and title, which it uses to create livestream
   - YouTube API returns the stream's key, iframe embed code, and URL, which the back-end sends to the front-end to be displayed for the user
   - User inputs URL and key into OBS Studio and starts livestreaming
   - After some seconds, user clicks on "GO LIVE" to see the iframe embedded
   - User can click on "End Livestream" to return to "not-live" state
For a more detailed explanation, refer to issue [that specifies the flow and features of Sprint 2](https://github.com/UF-CEN5035-2022SpringProject/GatorStore/issues/54)

4. Build the [API document](https://github.com/UF-CEN5035-2022SpringProject/GatorStore#api-document)
   We are using RESTful API with JSON format, and we names each APIs with there property of the feature.
   We have defined 
    - API URL
    - Header
    - Return Body
    - Error Code & Error Response

5. [Error code documents](https://github.com/UF-CEN5035-2022SpringProject/GatorStore#global-errorcode)
   Define proper error response for frontend to make the correct reaction.
   - Check issue [#62](https://github.com/UF-CEN5035-2022SpringProject/GatorStore/issues/62)
 
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

# GatorStore Seller Manage System
---
# How to Run
1. After starting the backend, then set the target backend api url correctly in settings.js
```
    apiHostURL: 'http://localhost:8080/api/',
    testApiHostURL: 'http://localhost:8080/test/api'
```
2. Start the application
```
npm install && npm start
```

- The application will be located at port 3000

# Routing To Pages Via /paths:
---
Routing will be done via the **react-router-dom library**, version 6.2.1.   
As of now, **index.js** holds the routes for all pages.  
To add a new route, add the following line below the existing routes in index.js (after importing the page's JS file):
```
<Route path="/path-to-new-page" element={<ComponentForNewPage/>} />
```

# Adding New Pages:
---
The JS files for a new page can be added to the **pages folder**. For each page, a CSS file of the same or similar name can be added to **styles folder**.  
There is also a **components folder** that holds the JS files for all recurring components, such as the header and footer.

# Styling Dependencies: Material UI
---
The projects styling dependencies are listed below:
* Material UI: @mui/material   |   version 5.3.1
* Material Icons: @material-ui/icons  |   version 5.3.1

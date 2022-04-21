# Sprint4
---
## What we have done in Final Sprint4
- [Final Project Description](https://github.com/UF-CEN5035-2022SpringProject/GatorStore/blob/main/sprint4.md#final-project-description)       
- Demo video functionality (3mins)
  - https://www.youtube.com/watch?v=umJ77oY6wDo 

- [Cypress test video](https://github.com/UF-CEN5035-2022SpringProject/GatorStore/blob/main/sprint4.md#cypress-test-video-1)
  - Buyer Testing 
  - Seller Testing
    
- [Backend unit test video](https://github.com/UF-CEN5035-2022SpringProject/GatorStore/blob/main/sprint4.md#backend-unit-test)
- [API Documentation](https://github.com/UF-CEN5035-2022SpringProject/GatorStore/blob/main/README.md)
- [Project board](https://github.com/orgs/UF-CEN5035-2022SpringProject/projects/1)
- Link to Sprint4 deliverables
  - 
- Frontend and backend team members
- Bonus points for public release

## Final Project Description
---
GatorStore is a livestreaming and e-commerce site that connects potential buyers to sellers via stream-enabled live communication, as well as not-live store browsing. A user can sign-in using Google Authentication as either a seller or buyer.     
By being a seller, a user is allowed to create unlimited stores, each of which they can populate with their products. Each store has the ability of going LIVE via the YouTube API, which allows sellers to name their livestreams and receive the necessary stream key from GatorStore. A record of a store's orders and livestream is kept for each store. These livestreams are designed to support sales, so sellers can display specific featured products in each. While the seller application is aimed at the seller creating and editing stores/products, they can also visit any of GatorStore's stores or products through their URLs.     
By being a buyer, a user can browse through the stores and browse through their inventory. From a product's page, they can view its details, store, and purchase it. During livestreams, buyers can interact with sellers through the embedded YouTube chat and browse through the featured items of the stream. A record of the seller's orders is kept for them.

## Achievements 
- 2 React web appication for over 10 pages component
- Develop over 20 Golang APIs
- Implement Google third party APIs
- Up to 130 issues solved
- Over 250 commits have been pushed

---
## Sprint Feature Description
**Front-end New Features:**
 - Created a new instructions page (at "/") for both buyer and seller projects
 - Created a new page for sellers to get the livestream history for each of their stores
 - Implemented page for sellers to see the order history for each of their stores
 - Fixed Bug described in [#124](https://github.com/UF-CEN5035-2022SpringProject/GatorStore/issues/124): Made it available for products to be bought from livestream (bug from past sprint 3)
 - Redesigned header and footer to more closely match UF theme
 - Replaced slider images from "/home" on buyer site with images relevant and specially made for GatorStore
 - Updated and finalized Cypress testing for buyer and seller projects

**Back-end New Features:**
 - Created [Store OrderList API](https://github.com/UF-CEN5035-2022SpringProject/GatorStore/issues/76#issuecomment-1076493741): Get the orders according to the store #120
 - Created [Get Live Orders](https://github.com/UF-CEN5035-2022SpringProject/GatorStore/issues/120#issuecomment-1097158961): Get the lives according to the store #120

### User tutorials
---
**Seller livestream tutorial**
Instructions for how to be a seller using GatorStore
- https://github.com/UF-CEN5035-2022SpringProject/GatorStore/blob/main/seller-frontend/seller-instructions.md

**Buyer livestream tutorial**
Instructions for buyers in GatorStore
- https://github.com/UF-CEN5035-2022SpringProject/GatorStore/blob/main/buyer-frontend/buyer-instructions.md

### Devloper tutorials
---
**Backend devloper tutorial**
- https://github.com/UF-CEN5035-2022SpringProject/GatorStore/blob/main/backend/backend-readme.md

**Frontend devloper tutorial**
- https://github.com/UF-CEN5035-2022SpringProject/GatorStore/blob/main/buyer-frontend/buyer-instructions.md
- https://github.com/UF-CEN5035-2022SpringProject/GatorStore/blob/main/seller-frontend/seller-instructions.md

---
## Demo video functionality

---
## Cypress test video 
Cypress testing for both web applications.

Buyer:

https://user-images.githubusercontent.com/40399062/164333895-316649c4-6227-4dd4-b2bb-1a63994b5f6b.mp4   

Seller:

https://user-images.githubusercontent.com/40399062/164333928-e97956af-8aa6-4d8e-a317-88b6ef85b1f7.mp4

---
## Backend Unit Test
We are using `Go test`.
Because part of the functionalities rely on Google services, we are not able to get 100% coverage(62%).

###  Backend unit test video 
https://user-images.githubusercontent.com/11768359/163632604-da2fd917-84c9-4c87-bc0f-cf2d0486ac5b.mp4

---
## API Documentation
https://github.com/UF-CEN5035-2022SpringProject/GatorStore/blob/main/README.md
---

---
## Sprint4 deliverables

---
## About the team
- Frontend team members
  - Sebastian Llerena
  - Yi-Ming Chang

- Backend team members
  - Hung-You Chou
  - Yi-Ming Chang

--- 
## Public release
Seller: 

Buyer: 



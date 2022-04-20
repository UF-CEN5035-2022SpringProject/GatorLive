// Test Home Page:
describe("Renders the home page", () => {
    it("Renders correctly and tests slider", () => {
        sessionStorage.clear() // try as a non-logged in user first

        cy.visit("/home")
        cy.get(".slider").should("exist")

        // Click right arrow and see that the second image is displayed
        cy.get('.right-arrow').click();
        cy.get('.active').find('img').should('have.attr', 'src').should('include', 'https://images.hindustantimes.com/tech/img/2021/10/02/960x540/Amazon_India_event_1633152363633_1633152377540.png')
    })

    it("Displays recommended stores API correctly", () => {
        cy.get(".MuiCardMedia-img").should("exist")
        // This is a class that only is used when mapping the recommended stores - meaning the API successfully returned them
    })
})

// Test a Store Page (gatorstore-1):
describe("Renders a store page correctly", () => {
    it("Renders store page information", () => {
        cy.visit("/store/gatorstore-1")
        cy.wait(4000) // wait for API to return store info from ID
        cy.get("#storename").contains("YiMing Fantastic Store") // check this store's Title
    })

    it('Check product page for product-1', function() {
        cy.visit('localhost:3001/store/gatorstore-1');
        cy.get(':nth-child(1) > .MuiPaper-root > .MuiCardMedia-root').should("exist");
        cy.get(':nth-child(1) > .MuiPaper-root > .MuiCardActions-root > .MuiButtonBase-root > .MuiButton-label').click();
    });
})

// Test a Product Page (product-1):
describe("Render a product page correctly", () => {
    it('Check product page info and button to go back to store, for product-1', () => {
        cy.get('.ProductKeyDetailColumn > h1').contains("YiMingSuperProduct-0"); // check title
        cy.get('.ProductStoreInfo > .MuiButton-root').click(); // go back to store
        
        // check that button led to store
        cy.url().then(($url) => {
            $url.includes("localhost:3001/store/gatorstore-1")
        })
    })

    it("Tests purchasing product without logging in", () => {
        cy.visit('localhost:3001/product/product-1');
        cy.get('.ProductKeyDetailColumn > div > .MuiButton-root').click(); // try to purchase product

        // system should display popup to tell them to sign first:
        cy.get('.stream-link-container > div > h2').contains("Please sign-in to purchase a product.");
        cy.get('.stream-link-container > div > .MuiButton-root').click(); // press OK
    })

    it("Tests purchasing product after logging in", () => {
        // set session data
        cy.window().then( win => {
            sessionStorage.setItem("user-name", "Yi-Ming Chang")
            sessionStorage.setItem("user-email", "ericblackking@gmail.com")
            sessionStorage.setItem("user-id", "11001")
            sessionStorage.setItem("user-jwtToken", "gst.R2F0b3JTdG9yZV9zaGViYXNxdWluZUBnbWFpbC5jb20xMTAwNA==_MjAyMi0wMi0yNVQyMTowMTo0OVo=")
        })

        cy.visit('localhost:3001/product/product-1');
        cy.get('.ProductKeyDetailColumn > div > .MuiButton-root').click(); // try to purchase product

        cy.wait(4000) // wait for API to send request

        cy.get('.stream-link-container > div > h2').contains('You have successfully purchased "YiMingSuperProduct-0"!'); // success message!
        cy.get('.stream-link-container > div > .MuiButton-root').click(); // press OK
    })
})

// Test the orders page:
describe("Render a Orders page correctly", () => {
    it("Tests going to orders page using Header", () => {
        cy.get('#basic-button').click(); // click on header user button
        cy.get('[href="/orders"]').click(); // click on orders

        cy.get('h1').contains("Your Orders"); // checks page title
    })
})

describe("Logs out successfully", () => {
    it("Test logging out", () => {
        cy.get('#basic-button').click(); // click on user button again
        cy.get('li.MuiMenuItem-root').click(); // logout

        // check you got re-directed to landing page:
        cy.url().then(($url) => {
            $url.includes("localhost:3001/")
        })
    })
})
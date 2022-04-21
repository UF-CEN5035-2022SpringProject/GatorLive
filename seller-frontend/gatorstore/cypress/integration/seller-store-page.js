describe("Creates Stream", () => {
    it("Renders correctly and creates stream", () => {
        // set session data
        cy.window().then( win => {
            sessionStorage.setItem("user-name", "Yi-Ming Chang")
            sessionStorage.setItem("user-email", "ericblackking@gmail.com")
            sessionStorage.setItem("user-id", "11001")
            sessionStorage.setItem("user-jwtToken", "gst.R2F0b3JTdG9yZV9zaGViYXNxdWluZUBnbWFpbC5jb20xMTAwNA==_MjAyMi0wMi0yNVQyMTowMTo0OVo=")
        })

        // go to store page
        cy.visit("/store/gatorstore-1");

        cy.wait(1000) // wait for API to send request

        // set stream information
        cy.get(':nth-child(1) > .MuiGrid-align-items-xs-center > .MuiGrid-justify-content-xs-flex-end > .MuiButton-root').click();
        cy.get('#titleField').type('Test Stream Title');
        cy.get('[style="text-align: center;"] > .MuiButton-containedPrimary').click();
        cy.get(':nth-child(1) > img').click();
        cy.get(':nth-child(2) > img').click();
        
        // Go live and see stream iframe
         cy.get('.stream-link-container > div > .MuiButton-root').click();
         cy.get('[style="text-align: center;"] > .MuiButton-root').click();

        // Check for UI's responsiveness to going to the "live" state
        //cy.get(".streamChat").should("exist");
    })
})
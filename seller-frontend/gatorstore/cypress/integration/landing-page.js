// Test image slider:
describe("Renders the home page", () => {
    it("Renders correctly and tests slider", () => {
        cy.visit("/landingpage")
        cy.get(".slider").should("exist")

        // Click right arrow and see that the second image is displayed
        cy.get('.right-arrow').click();
        cy.get('.active').find('img').should('have.attr', 'src').should('include', 'https://images.hindustantimes.com/tech/img/2021/10/02/960x540/Amazon_India_event_1633152363633_1633152377540.png')
    })
})

// Testing Google Login
//  1. Attempt 1: Programmatically calling Google's API via a POST (code for loginByGoogleApi() on commands.js)
/*describe('Google Login Test', function () {
  beforeEach(function () {
    cy.loginByGoogleApi()
  })

  it('Shows successful login ', function () {
    cy.get(".accountButton").should("exist"); // checks that the container containing the user's email is there instead of the "Sign In" button 
  })
})*/

//  2. Attempt 2: Using Cypress's "cypress-social-logins" library
describe('Login', () => {
    it('Login through Google', () => {
      const username = Cypress.env('googleSocialLoginUsername')
      const password = Cypress.env('googleSocialLoginPassword')
      const loginUrl = Cypress.env('loginUrl')
      const cookieName = "cookie"
      const socialLoginOptions = {
        username: username,
        password: password,
        loginUrl: loginUrl,
        headless: false,
        logs: false,
        loginSelector: 'a[id="loginButton"]',
        postLoginSelector: '.slider'
      }

      cy.clearCookies()
  
      return cy.task('GoogleSocialLogin', socialLoginOptions).then(({cookies}) => {
        cy.clearCookies()
  
        const cookie = cookies.filter(cookie => cookie.name === cookieName).pop()
        if (cookie) {
          cy.setCookie(cookie.name, cookie.value, {
            domain: cookie.domain,
            expiry: cookie.expires,
            httpOnly: cookie.httpOnly,
            path: cookie.path,
            secure: cookie.secure
          })
  
          Cypress.Cookies.defaults({
            preserve: cookieName
          })
        }
      })
    })
  })
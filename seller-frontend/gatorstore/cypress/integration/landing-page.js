describe("Renders the home page", () => {
    it("Renders correctly and tests slider", () => {
        cy.visit("/landingpage")
        cy.get(".slider").should("exist")

        /* ==== Click right arrow and see that the second image is displayed ==== */
        cy.get('.right-arrow').click();
        cy.get('.active').find('img').should('have.attr', 'src').should('include', 'https://images.hindustantimes.com/tech/img/2021/10/02/960x540/Amazon_India_event_1633152363633_1633152377540.png')
    })
})

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
        headless: true,
        logs: false,
        loginSelector: 'a[id="loginButton"]',
        postLoginSelector: '.slider'
      }
  
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
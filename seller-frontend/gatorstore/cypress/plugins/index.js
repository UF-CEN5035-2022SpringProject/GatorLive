/// <reference types="cypress" />
// ***********************************************************
// This example plugins/index.js can be used to load plugins
//
// You can change the location of this file or turn off loading
// the plugins file with the 'pluginsFile' configuration option.
//
// You can read more here:
// https://on.cypress.io/plugins-guide
// ***********************************************************

// This function is called when a project is opened or re-opened (e.g. due to
// the project's config changing)

//const {GoogleSocialLogin} = require('cypress-social-logins').plugins


// .env
REACT_APP_GOOGLE_CLIENTID = 'your-client-id'
REACT_APP_GOOGLE_CLIENT_SECRET = 'your-client-secret'
GOOGLE_REFRESH_TOKEN = 'your-refresh-token'

//dotenv.config()


/**
 * @type {Cypress.PluginConfig}
 */
// eslint-disable-next-line no-unused-vars
module.exports = (on, config) => {
  // `on` is used to hook into various events Cypress emits
  // `config` is the resolved Cypress config
  /*on('task', {
    GoogleSocialLogin: GoogleSocialLogin
  })*/

  
  config.env.googleRefreshToken = REACT_APP_GOOGLE_CLIENTID 
  config.env.googleClientId = REACT_APP_GOOGLE_CLIENTID
  config.env.googleClientSecret = REACT_APP_GOOGLE_CLIENT_SECRET 

  return config
}
const domain = 'seller.gatorstore.org'
const apiDomain = 'localhost'
export default {
    HeaderText: {
        'fontFamily': 'Titillium Web',
        'fontSize': 16
    },
    Colors: {
        'mainColor': '#e10600',
        'subColor': '#fff',
        'thirdColor': '#15151e',
        'forthColor': '#38383f'
    },
    Font: {
        'major': 'Titillium Web',
        'secondary': 'Audiowide'
    },
    apiHostURL: `http://${apiDomain}:8080/api/`,
    testApiHostURL: `http://${apiDomain}:8080/test/api`,
    googleLoginRedirectURL: `http%3A%2F%2F${domain}%3A3000%2Flogin`,
    applicationRootURL: `http://${domain}:3000`,
    buyerAppURL: 'https://buyer.gatorstore.org/',
    domain: domain
}
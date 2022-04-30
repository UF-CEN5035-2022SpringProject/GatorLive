const domain = 'buyer.gatorstore.org'
export default {
    HeaderText: {
        'fontFamily': 'Ubuntu',
        'fontSize': 16
    },
    Colors: {
        'mainColor': '#e10600',
        'subColor': '#fff',
        'thirdColor': '#15151e',
        'forthColor': '#38383f'
    },
    Font: {
        'major': 'Ubuntu',
        'secondary': 'Audiowide'
    },
    googleLoginRedirectURL: `https%3A%2F%2F${domain}%2Flogin`,
    applicationRootURL: `https://${domain}`,
    applicationPort: 3001,
    applicationHost: domain,
    apiHostURL: `https://${domain}/api/`,
    testApiHostURL: `https://${domain}/test/api`,
    sellerAppURL: 'https://seller.gatorstore.org'
}
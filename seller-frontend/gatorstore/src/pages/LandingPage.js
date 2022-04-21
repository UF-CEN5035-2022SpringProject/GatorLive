import React, { useRef } from 'react';
import Header from '../components/Header.js';
import Footer from '../components/Footer';
import Paper from '@mui/material/Paper';
import { styled } from '@mui/material/styles';
import GoogleIcon from '@mui/icons-material/Google';
import settings from '../settings.js';

const Item = styled(Paper)(({ theme }) => ({
    ...theme.typography.body2,
    textAlign: 'center',
    color: theme.palette.text.secondary,
    height: 60,
    lineHeight: '60px',
}));

const CardElevation = 24;

function GoogleButton() { // for when user is NOT signed in
    return (<div className="g-signin">
      <a id="loginButton" style={{textDecoration: 'none'}} href={`https://accounts.google.com/o/oauth2/auth?access_type=offline&client_id=138444517704-gg6649ok973letdlh55bpte8bna7721o.apps.googleusercontent.com&redirect_uri=${settings.googleLoginRedirectURL}&response_type=code&scope=https%3A%2F%2Fwww.googleapis.com%2Fauth%2Fyoutube+https%3A%2F%2Fwww.googleapis.com%2Fauth%2Fuserinfo.email+https%3A%2F%2Fwww.googleapis.com%2Fauth%2Fuserinfo.profile&state=state`}>
        <button className='login' style={{marginLeft: 0}}>
          <GoogleIcon style={{verticalAlign: 'middle', marginRight: 10}}/> Sign-In / Sign-Up 
        </button>
      </a>
    </div>);
  }
  

function RedirectToBuyerButton() { // for when user is NOT signed in
    return (<div className="g-signin">
        <a id="redirectToBuyerButton" style={{textDecoration: 'none'}} href={settings.buyerAppURL}>
        <button className='login' style={{marginLeft: 0}}>
            <GoogleIcon style={{verticalAlign: 'middle', marginRight: 10}}/> Browse the Stores 
        </button>
        </a>
    </div>);
}

function LandingPage() {
    const introOpenStore = useRef(null);
    const introProductDisplay = useRef(null);
    const introLivePromotions = useRef(null);
    const introOrderFeature = useRef(null);
    const introTeam = useRef(null);

    const toOpenStoreScroll = () => introOpenStore.current.scrollIntoView();
    const toProductDisplayScroll = () => introProductDisplay.current.scrollIntoView();
    const toLivePromotionsScroll = () => introLivePromotions.current.scrollIntoView();
    const toOrdersFeature = () => introOrderFeature.current.scrollIntoView();
    const toTeamScroll = () => introTeam.current.scrollIntoView();

    return( 
        <div>
            <div >
                <Header/>
            </div>
            <div style={{ 
                display: 'flex',
                height: 400, 
                width: '100%', 
                marginTop: 2, 
                justifyContent: 'center',
                backgroundColor: 'blue', /* For browsers that do not support gradients */
                backgroundImage: 'linear-gradient(to bottom right,red 0%,rgb(224, 129, 46) 100%)'
            }}>
                <div style={{
                    display: 'flex', 
                    height: '100%', 
                    width: '100%', 
                    flexDirection: 'column',
                }}> 
                    <div style={{marginTop: 30}}/>
                    <div className='instruction-title' style={{display: 'flex', justifyContent: 'center'}}>
                        Start Your Shopping Channel
                    </div>
                    <div style={{marginTop: 15}}/>
                    <div className='instruction-content' style={{display: 'flex', justifyContent: 'center'}}>
                        We are not simply a E-commerce platform. <br/>
                        Start your Shopping Channel to make your products sell dramatically fast! <br/>
                        As a Seller, join GatorStore! <br/>
                    </div>
                    <div style={{marginTop: 70}}/>
                    <div style={{width: '100%', height: 70, display: 'flex', flexDirection: 'row', justifyContent: 'space-evenly'}}>
                        <div className='instruction-btn' onClick={toOpenStoreScroll} style={{height: '100%', width:300}}>
                            Open up Stores
                        </div>
                        <div className='instruction-btn' onClick={toProductDisplayScroll} style={{height: '100%', width:300}}>
                            Product display
                        </div>
                        <div className='instruction-btn' onClick={toLivePromotionsScroll} style={{height: '100%', width:300}}>
                            Live Promotions 
                        </div>
                        <div className='instruction-btn' onClick={toOrdersFeature} style={{height: '100%', width:300}}>
                            Order Management 
                        </div>
                        <div className='instruction-btn' onClick={toTeamScroll} style={{height: '100%', width:300}}>
                            Meet the Team
                        </div>
                    </div>
                </div>
            </div>
            <div style={{ 
                display: 'flex',
                height: 600, 
                width: '100%', 
                marginTop: 10, 
                flexDirection: 'row',
                justifyContent: 'center',
                alignItems: 'center'
            }}>
                <div style={{width: '2%'}}/>
                <div 
                    ref={introOpenStore} 
                    style={{
                    display: 'flex', 
                    height: '100%', 
                    width: '46%', 
                    flexDirection: 'column',
                    justifyContent: 'center',
                    textAlign: 'right' 
                }}> 
                    <div className='block-title'>
                        Open up Stores
                    </div><br/> 
                    <div className='block-content'>
                        Create your own store for customer to online-shopping easily. <br/> <br/> 
                        Manage multiple store in the same platform efficiently. 
                    </div>
                </div>
                <div style={{width: '2%'}}/>
                <div style={{width: 2, height: 300, backgroundColor: '#FA4616'}}/>
                <div style={{width: '2%'}}/>
                <img 
                    alt='storeList' 
                    src='/storeListImage.png'
                    className='block-img'
                    style={{
                        height: 500, 
                        width: '46%', 
                        marginTop: 106.5, 
                        borderRadius: 15
                    }}
                />  
                <div style={{width: '2%'}}/>
            </div>
            <div
                ref={introProductDisplay} 
                style={{ 
                display: 'flex',
                height: 600, 
                width: '100%', 
                marginTop: 10, 
                flexDirection: 'row',
                justifyContent: 'center',
                alignItems: 'center'
            }}>
                <div style={{width: '2%'}}/>
                <img 
                    alt='storeProducts' 
                    src='/storeProductImage.png'
                    className='block-img'
                    style={{
                        height: 500, 
                        width: '46%', 
                        marginTop: 106.5, 
                        borderRadius: 15
                    }}
                /> 
                <div style={{width: '2%'}}/>
                <div style={{width: 2, height: 300, backgroundColor: '#FA4616'}}/>
                <div style={{width: '2%'}}/>
                <div style={{
                    display: 'flex', 
                    height: '100%', 
                    width: '46%', 
                    flexDirection: 'column',
                    justifyContent: 'center',
                    textAlign: 'left' 
                }}> 
                    <div className='block-title'>
                        Product display 
                    </div><br/> 
                    <div className='block-content'>
                        The Store Page is the hub for live user interactions and browsing through the products. <br/> <br/> 
                        You can visit each of the products to buy them by clicking on "View More" beneath each. 
                    </div>
                </div> 
                <div style={{width: '2%'}}/>
            </div>
            <div
                ref={introLivePromotions}
                style={{ 
                display: 'flex',
                height: 600, 
                width: '100%', 
                marginTop: 10, 
                flexDirection: 'row',
                justifyContent: 'center',
                alignItems: 'center'
            }}>
                <div style={{width: '2%'}}/>
                <div style={{
                    display: 'flex', 
                    height: '100%', 
                    width: '46%', 
                    flexDirection: 'column',
                    justifyContent: 'center',
                    textAlign: 'right' 
                }}> 
                    <div className='block-title'>
                        Promote products by Live 
                    </div><br/> 
                    <div className='block-content'>
                        Each store is equipped with the ability to go LIVE. Choose some "featured products" you wish to showcase for them! <br/> <br/> 
                        You can click on the "START LIVESTREAM" button to start your own YouTube embedded livestream.
                    </div>
                </div>
                <div style={{width: '2%'}}/>
                <div style={{width: 2, height: 300, backgroundColor: '#FA4616'}}/>
                <div style={{width: '2%'}}/>
                <img 
                    alt='storeLiveStream' 
                    src='/liveStreamImage.png'
                    className='block-img'
                    style={{
                        height: 500, 
                        width: '46%', 
                        marginTop: 106.5, 
                        borderRadius: 15
                    }}
                />  
                <div style={{width: '2%'}}/>
            </div>
            <div
                style={{ 
                display: 'flex',
                height: 600, 
                width: '100%', 
                marginTop: 10, 
                flexDirection: 'row',
                justifyContent: 'center',
                alignItems: 'center'
            }}>
                <div style={{width: '2%'}}/>
                <img 
                    alt='liveObs' 
                    src='/liveObs.gif'
                    className='block-img'
                    style={{
                        height: 500, 
                        width: '46%', 
                        marginTop: 106.5, 
                        borderRadius: 15
                    }}
                />
                <div style={{width: '2%'}}/>
                <div style={{width: 2, height: 300, backgroundColor: '#FA4616'}}/>
                <div style={{width: '2%'}}/>
                <div style={{
                    display: 'flex', 
                    height: '100%', 
                    width: '46%', 
                    flexDirection: 'column',
                    justifyContent: 'center',
                    textAlign: 'left' 
                }}> 
                    <div className='block-title'>
                        Stream by OBS 
                    </div><br/> 
                    <div className='block-content'>
                        Once you have named your livestream and chosen its "featured products", you will receive the key and URL needed to start streaming. <br/> <br/> 
                        By inputting those into OBS, you will be able to control your stream from your own system.
                    </div>
                </div>
                <div style={{width: '2%'}}/>
            </div>
            <div
                style={{ 
                display: 'flex',
                height: 600, 
                width: '100%', 
                marginTop: 10, 
                flexDirection: 'row',
                justifyContent: 'center',
                alignItems: 'center'
            }}>
                <div style={{width: '2%'}}/>
                <div style={{
                    display: 'flex', 
                    height: '100%', 
                    width: '46%', 
                    flexDirection: 'column',
                    justifyContent: 'center',
                    textAlign: 'right' 
                }}> 
                    <div className='block-title'>
                        Live History with Orders 
                    </div><br/> 
                    <div className='block-content'>
                        Live History page shows how much orders come from your livestream <br/> <br/> 
                        To visit it, place "/livestreams" after your store's name. It will list all of the orders made to featured products during it.
                    </div>
                </div>
                <div style={{width: '2%'}}/>
                <div style={{width: 2, height: 300, backgroundColor: '#FA4616'}}/>
                <div style={{width: '2%'}}/>
                <img 
                    alt='liveHistoryListImage' 
                    src='/liveHistoryListImage.png'
                    className='block-img'
                    style={{
                        height: 500, 
                        width: '46%', 
                        marginTop: 106.5, 
                        borderRadius: 15
                    }}
                />  
                <div style={{width: '2%'}}/>
            </div>
            <div
                ref={introOrderFeature} 
                style={{ 
                display: 'flex',
                height: 600, 
                width: '100%', 
                marginTop: 10, 
                flexDirection: 'row',
                justifyContent: 'center',
                alignItems: 'center'
            }}>
                <div style={{width: '2%'}}/>
                <img 
                    alt='orderHistory' 
                    src='/orderHistoryImage.png'
                    className='block-img'
                    style={{
                        height: 500, 
                        width: '46%', 
                        marginTop: 106.5, 
                        borderRadius: 15
                    }}
                /> 
                <div style={{width: '2%'}}/>
                <div style={{width: 2, height: 300, backgroundColor: '#FA4616'}}/>
                <div style={{width: '2%'}}/>
                <div style={{
                    display: 'flex', 
                    height: '100%', 
                    width: '46%', 
                    flexDirection: 'column',
                    justifyContent: 'center',
                    textAlign: 'left' 
                }}> 
                    <div className='block-title'>
                        Order Management 
                    </div><br/> 
                    <div className='block-content'>
                        Additionally, the order management page shows orders from one of your stores. <br/> <br/> 
                        To view it, place "/orders" after your store's name.
                    </div>
                </div> 
                <div style={{width: '2%'}}/>
            </div>
            <div style={{
                display: "flex", 
                height: '600px', 
                width: '100%', 
                marginTop: 100, 
                justifyContent: 'center',
                alignItems: 'center',
                backgroundImage: "url(/tower-login.jpeg)",
                backgroundRepeat: "no-repeat",
                backgroundPosition: "center",
                backgroundSize: "cover",
            }}>
                <div className='instruction-card-root instruction-card-outlined instruction-card-rounded' style={{
                    height: '45%', 
                    width: '30%',
                    display: 'flex', 
                    flexDirection: 'column',
                    alignItems: 'center'
                }}>
                    <div className='instruction-title' style={{color: 'rgba(0, 0, 0, 0.75)', fontSize: 32, marginTop: 15, marginBottom: 15}}>Join as Seller</div>
                    <GoogleButton />
                    <div style={{width: '80%', height: 2, backgroundColor: 'rgba(0, 0, 0, 0.4)', marginTop: 15}}></div>
                    <div className='instruction-title' style={{color: 'rgba(0, 0, 0, 0.75)', fontSize: 32, marginTop: 15, marginBottom: 15}}>Buy on GatorStore</div>
                    <RedirectToBuyerButton />
                </div>
            </div>
            <div ref={introTeam} style={{marginTop: 30}}>
                <Footer/>
            </div>

        </div>
    );
}

export default LandingPage;
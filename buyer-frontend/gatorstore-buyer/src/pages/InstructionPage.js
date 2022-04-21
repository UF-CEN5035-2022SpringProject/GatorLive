import React, { useRef } from 'react';
import Header from '../components/Header.js';
import Footer from '../components/Footer';
import Paper from '@mui/material/Paper';
import { styled } from '@mui/material/styles';
import GoogleIcon from '@mui/icons-material/Google';
import settings from '../settings.js';
import { textAlign } from '@mui/system';

import instruction1 from '../images/buyerInstruction1.png';
import instruction2 from '../images/buyerInstruction2.png';
import instruction3 from '../images/buyerInstruction3.png';
import instruction4 from '../images/buyerInstruction4.png';

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

function RedirectToSellerButton() { // for when user is NOT signed in
    return (<div className="g-signin">
        <a id="redirectToSellerButton" style={{textDecoration: 'none'}} href={settings.sellerAppURL}>
        <button className='login' style={{marginLeft: 0}}>
            <GoogleIcon style={{verticalAlign: 'middle', marginRight: 10}}/> Join GatorStore as a Seller 
        </button>
        </a>
    </div>);
}

function InstructionPage() {
    const introGatorStore = useRef(null);
    const introLiveShop = useRef(null);
    const introFollowOrders = useRef(null);
    const introTeam = useRef(null);

    const toGatorStoreScroll = () => introGatorStore.current.scrollIntoView();
    const toLiveShopScroll = () => introLiveShop.current.scrollIntoView();
    const toFollowOrdersScroll = () => introFollowOrders.current.scrollIntoView();
    const toTeamScroll = () => introTeam.current.scrollIntoView();

    return( 
        <div>
            <div >
                <Header />
            </div>
            <div style={{ 
                display: 'flex',
                height: 400, 
                width: '100%', 
                marginTop: 2, 
                justifyContent: 'center',
                backgroundColor: 'blue', /* For browsers that do not support gradients */
                backgroundImage: 'linear-gradient(to bottom right,rgb(2,3,129) 0%,rgb(40,116,252) 100%)'
            }}>
                <div style={{
                    display: 'flex', 
                    height: '100%', 
                    width: '100%', 
                    flexDirection: 'column',
                }}> 
                    <div style={{marginTop: 30}}/>
                    <div className='instruction-title' style={{display: 'flex', justifyContent: 'center'}}>
                        What is GatorStore?
                    </div>
                    <div style={{marginTop: 15}}/>
                    <div className='instruction-content' style={{display: 'flex', justifyContent: 'center'}}>
                        By reading the name we are not simply a E-commerce platform. <br/>
                        Watch our Seller's Live Streams to find the product you want! <br/>
                        Get special discount in the Live Rooms! <br/>
                    </div>
                    <div style={{marginTop: 70}}/>
                    <div style={{width: '100%', height: 70, display: 'flex', flexDirection: 'row', justifyContent: 'space-evenly'}}>
                        <div className='instruction-btn' onClick={toGatorStoreScroll} style={{height: '100%', width:300}}>
                            Join GatorStore
                        </div>
                        <div className='instruction-btn' onClick={toLiveShopScroll} style={{height: '100%', width:300}}>
                            Live Shopping
                        </div>
                        <div className='instruction-btn' onClick={toFollowOrdersScroll} style={{height: '100%', width:300}}>
                            Purchase a Product
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
                    ref={introGatorStore} 
                    style={{
                    display: 'flex', 
                    height: '100%', 
                    width: '46%', 
                    flexDirection: 'column',
                    justifyContent: 'center',
                    textAlign: 'right' 
                }}> 
                    <div className='block-title'>
                            Join GatorStore
                    </div><br/>
                    <div className='block-content'>
                        First, become a user with your Google account by clicking on the "Sign-In" button on the top-right corner. <br/> <br/>
                        You will be prompted to grant us permission to sign you in using your Google account.
                    </div>
                </div>
                <div style={{width: '2%'}}/>
                <div style={{width: 2, height: 300, backgroundColor: '#FA4616'}}/>
                <div style={{width: '2%'}}/>
                <img 
                    alt='storeProducts' 
                    src={instruction1}
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
                ref={introLiveShop} 
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
                    src={instruction2}
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
                        Live Shopping
                    </div><br/> 
                    <div className='block-content'>
                        Click on one of the recommended stores or on "My Favorite Store" in the top-right header menu!<br/> <br/> 
                        Whether it is offline or live, you can browse the store's inventory at the page's bottom.<br/> <br/> 
                        When live, the top will expand to include the livestream, chat, and featured products - which are the products the stream is showcasing.
                    </div>
                </div> 
                <div style={{width: '2%'}}/>
            </div>
            <div
                ref={introFollowOrders}
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
                        Purchase a Product
                    </div><br/> 
                    <div className='block-content'>
                        By clicking "View More" on a product card, you can reach its page.<br/> <br/> 
                        This page contains the product's image, price, description, and stock status. If the button shown here is blue, you can buy it!
                    </div>
                </div>
                <div style={{width: '2%'}}/>
                <div style={{width: 2, height: 300, backgroundColor: '#FA4616'}}/>
                <div style={{width: '2%'}}/>
                <img 
                    alt='storeProducts' 
                    src={instruction3}
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
                    alt='storeProducts' 
                    src={instruction4}
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
                        Follow your Orders
                    </div><br/> 
                    <div className='block-content'>
                        Using the top-right header menu, you can visit your order history page. <br/> <br/> 
                        In it, order details and a link to each of the products can be found.
                    </div>
                </div> 
                <div style={{width: '2%'}}/>
            </div>

            {/* Start Browsing/Become a Seller Block Below: */}

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
                    <div className='instruction-title' style={{color: 'rgba(0, 0, 0, 0.75)', fontSize: 32, marginTop: 15, marginBottom: 15}}>Start Browsing</div>
                    <GoogleButton />
                    <div style={{width: '80%', height: 2, backgroundColor: 'rgba(0, 0, 0, 0.4)', marginTop: 15}}></div>
                    <div className='instruction-title' style={{color: 'rgba(0, 0, 0, 0.75)', fontSize: 32, marginTop: 15, marginBottom: 15}}>Become a Seller</div>
                    <RedirectToSellerButton />
                </div>
            </div>

            <div ref={introTeam} style={{marginTop: 30}}>
                <Footer/>
            </div>

        </div>
    );
}

export default InstructionPage;
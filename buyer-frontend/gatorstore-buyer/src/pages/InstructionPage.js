import React from 'react';
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
      <a id="loginButton" style={{textDecoration: 'none'}} href={`https://accounts.google.com/o/oauth2/auth?access_type=offline&client_id=138444517704-gg6649ok973letdlh55bpte8bna7721o.apps.googleusercontent.com&redirect_uri=http%3A%2F%2Flocalhost%3A${settings.applicationPort}%2Flogin&response_type=code&scope=https%3A%2F%2Fwww.googleapis.com%2Fauth%2Fyoutube+https%3A%2F%2Fwww.googleapis.com%2Fauth%2Fuserinfo.email+https%3A%2F%2Fwww.googleapis.com%2Fauth%2Fuserinfo.profile&state=state`}>
        <button className='login' style={{marginLeft: 0}}>
          <GoogleIcon style={{verticalAlign: 'middle', marginRight: 10}}/> Sign-In / Sign-Up 
        </button>
      </a>
    </div>);
  }

function RedirectToSellerButton() { // for when user is NOT signed in
    return (<div className="g-signin">
        <a id="redirectToSellerButton" style={{textDecoration: 'none'}} href={`http://${settings.applicationHost}:3000`}>
        <button className='login' style={{marginLeft: 0}}>
            <GoogleIcon style={{verticalAlign: 'middle', marginRight: 10}}/> Join GatorStore as a Seller 
        </button>
        </a>
    </div>);
}

function InstructionPage() {
    return( 
        <div>
            <div >
                <Header/>
            </div>
            <div style={{ 
                display: 'flex',
                height: 400, 
                width: '100%', 
                marginTop: 10, 
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
                        <div className='instruction-btn' style={{height: '100%', width:300}}>
                            Browse the Stores
                        </div>
                        <div className='instruction-btn' style={{height: '100%', width:300}}>
                            Live Shopping
                        </div>
                        <div className='instruction-btn' style={{height: '100%', width:300}}>
                            Follow Your Orders
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
            }}>
                <div style={{width: '2%'}}/>
                <div style={{
                    display: 'flex', 
                    height: '100%', 
                    width: '47%', 
                    flexDirection: 'column',
                }}> 
                    <div style={{marginTop: 30}}/>
                    <div className='block-title' style={{display: 'flex', justifyContent: 'center'}}>
                        Browse the Stores
                    </div>
                    <Item 
                        elevation={CardElevation}
                        style={{marginTop: 30, height: 500}}
                    >
                        <div className='block-content'>
                            The Store Page is the hub for live user interactions and browsing through the products. You can visit each of the products to buy them by clicking on "View More" beneath each. If the store is live, the YouTube livestream and chat will automatically be embedded in the top of the page, alongside its featured products. Even when not livestreaming, you will be able to browse through a store's products on the lower portion of the page.
                        </div>
                    </Item>
                </div>
                <div style={{width: '2%'}}/>
                <img 
                    alt='storeProducts' 
                    src='/storeProductImage.png'
                    style={{height: 500, width: '47%', marginTop: 106.5}}
                />  
                <div style={{width: '2%'}}/>
            </div>
            <div style={{ 
                display: 'flex',
                height: 600, 
                width: '100%', 
                marginTop: 50, 
                flexDirection: 'row',
                justifyContent: 'center',
            }}>
                <div style={{width: '2%'}}/>
                <img 
                    alt='storeProducts' 
                    src='/storeProductImage.png'
                    style={{height: 500, width: '47%', marginTop: 106.5}}
                />
                <div style={{width: '2%'}}/>
                <div style={{
                    display: 'flex', 
                    height: '100%', 
                    width: '47%', 
                    flexDirection: 'column',
                }}> 
                    <div style={{marginTop: 30}}/>
                    <div className='block-title' style={{display: 'flex', justifyContent: 'center'}}>
                        Live Shopping
                    </div>
                    <Item 
                        elevation={CardElevation}
                        style={{marginTop: 30, height: 500}}
                    >
                        <div className='block-content'>
                            When a seller's store goes live, its page will showcase the livestream and chat! <br/>Keep up with a store to shop through live seller/buyer communication!
                        </div>
                    </Item>
                </div>
                <div style={{width: '2%'}}/>
            </div>
            <div style={{ 
                display: 'flex',
                height: 600, 
                width: '100%', 
                marginTop: 50, 
                flexDirection: 'row',
                justifyContent: 'center',
            }}>
                <div style={{width: '2%'}}/>
                <div style={{
                    display: 'flex', 
                    height: '100%', 
                    width: '47%', 
                    flexDirection: 'column',
                }}> 
                    <div style={{marginTop: 30}}/>
                    <div className='block-title' style={{display: 'flex', justifyContent: 'center'}}>
                        Follow Your Orders
                    </div>
                    <Item 
                        elevation={CardElevation}
                        style={{marginTop: 30, height: 500}}
                    >
                        <div className='block-content'>
                            Visit the "My Orders" page through the header (top-right button) to see your order history! <br/>Details such as their quantities, store, title, and product page can be found here.
                        </div>
                    </Item>
                </div>
                <div style={{width: '2%'}}/>
                <img 
                    alt='storeProducts' 
                    src='/storeProductImage.png'
                    style={{height: 500, width: '47%', marginTop: 106.5}}
                />  
                <div style={{width: '2%'}}/>
            </div>
            <div style={{
                display: "flex", 
                height: '600px', 
                width: '100%', 
                marginTop: 30, 
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
                    <div className='instruction-title' style={{color: 'rgba(0, 0, 0, 0.75)', fontSize: 32, marginTop: 15, marginBottom: 15}}>Befome a Seller</div>
                    <RedirectToSellerButton />
                </div>
            </div>
            <div>
                <Footer/>
            </div>

        </div>
    );
}

export default InstructionPage;
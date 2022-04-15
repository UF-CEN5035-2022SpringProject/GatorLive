import React from 'react';
import Header from '../components/Header.js';
import Footer from '../components/Footer';
import Paper from '@mui/material/Paper';
import { createTheme, ThemeProvider, styled } from '@mui/material/styles';

const Item = styled(Paper)(({ theme }) => ({
    ...theme.typography.body2,
    textAlign: 'center',
    color: theme.palette.text.secondary,
    height: 60,
    lineHeight: '60px',
}));

const lightTheme = createTheme({ palette: { mode: 'light' } });
const CardElevation = 24;

function InstructionPage() {
    return( 
        <div>
            <div >
                <Header/>
            </div>
            <div style={{ 
                display: 'flex',
                height: 600, 
                width: '100%', 
                marginTop: 10, 
                justifyContent: 'center',
                backgroundColor: 'blue', /* For browsers that do not support gradients */
                backgroundImage: 'linear-gradient(to bottom right,rgb(2,3,129) 0%,rgb(40,116,252) 100%)'
            }}>
                <div style={{
                    display: 'flex', 
                    height: '100%', 
                    width: '50%', 
                    flexDirection: 'column',
                }}> 
                    <div style={{marginTop: 30}}/>
                    <div className='instruction-title' style={{display: 'flex', justifyContent: 'center'}}>
                        What is GatorStore?
                    </div>
                    <div style={{marginTop: 15}}/>
                    <div className='instruction-content'>
                        By reading the name we are not simply a E-commerce platform. <br/>
                        We provide user using live streams to advise their products and get more attentions. <br/>
                        Boost the profit by using this spectacular feature! <br/>
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
                        Watch Live Shopping
                    </div>
                    <Item 
                        className='block-content' 
                        elevation={CardElevation}
                        style={{marginTop: 30, height: 500}}
                    >
                        TEST
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
                        className='block-content' 
                        elevation={CardElevation}
                        style={{marginTop: 30, height: 500}}
                    >
                        TEST
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
                alignContent: 'center',
                backgroundImage: "url(/tower-login.jpeg)",
                backgroundRepeat: "no-repeat",
                backgroundPosition: "center",
                backgroundSize: "cover",
            }}>
                <div style={{
                    display: "flex", 
                    height: '600px', 
                    width: '800px'
                }}>
                </div>
            </div>
            <div>
                <Footer/>
            </div>

        </div>
    );
}

export default InstructionPage;
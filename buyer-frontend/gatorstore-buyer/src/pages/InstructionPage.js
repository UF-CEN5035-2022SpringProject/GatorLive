import React, {useState, useEffect} from 'react';
import Header from '../components/Header.js';
import Footer from '../components/Footer';
import { alignProperty } from '@mui/material/styles/cssUtils';


function InstructionPage() {
    return( 
        <div>
            <div >
                <Header/>
            </div>
            <div style={{
                display: "flex", 
                height: '600px', 
                width: '100%', 
                marginTop: 30, 
                justifyContent: 'center'
            }}>
                <div className='instruction-title' style={{fontSize: 40}}>
                    What is GatorStore?
                </div>
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
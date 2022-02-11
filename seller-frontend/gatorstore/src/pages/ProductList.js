import React, {useState, useEffect} from 'react';
import Header from '../components/Header';
import Footer from '../components/Footer';
import Button from '@mui/material/Button';
import {Grid} from "@material-ui/core";
import Productcontent from '../components/ProductContent';

import CircleIcon from '@mui/icons-material/Circle';
import ContentCopyIcon from '@mui/icons-material/ContentCopy';
import '../styles/sellerStorePage.css';

function ProductList() {
  var storeObject;
  CheckStoreObject(); // initial check

  // Checks the status of this store's object in the store API:
  function CheckStoreObject() { 
    // storeObject = fetch and stuff from API
    // make sure to async and wait before rendering out this page
    alert("h");
  }

  // Run check-object-function every 60 seconds:
  useEffect(() => {
    const interval = setInterval(() => {
      CheckStoreObject();
    }, 60000);
  
    return () => clearInterval(interval); // This represents the unmount function, in which you need to clear your interval to prevent memory leaks.
  }, [])

  // Overlay hook:
  const [overlayDisplayed, ChangeOverlayDisplayed] = useState("none");

  useEffect(() => {
    document.body.style.overflowY = (overlayDisplayed != "none") ? "hidden" : "scroll";
  }, [overlayDisplayed]);

  // Overlay displaying url and key for the new stream:
  function StreamCreatedOverlay() {
    return(
      <div class="go-live-overlay" style={{display: overlayDisplayed}}>
        <div class="transparentBG"/>
        <div class="stream-link-container">
          <h1>Stream Created</h1>
          <p>Here are your stream's url:</p>

          <div class="stream-url-box">
            <p>https://support.google.com/youtube/answer/9854503?hl=en</p>
            <Button variant="contained" color="secondary" onClick={() => {}}><ContentCopyIcon/></Button>
          </div>

          <div style={{textAlign: "center"}}>
            <Button variant="contained" color="primary" onClick={() => {ChangeOverlayDisplayed("none") }} size="large">Continue</Button>
          </div>
        </div>
      </div>
    );
  }

  function GoLive() {
    // call API

    ChangeOverlayDisplayed("block");
  }

  function LiveInfoBar() {
    return(
      <Grid container spacing={0} justifyContent="center" alignItems="center" direction='row' style={{marginBottom: 20}}>
        <Grid item md={4} container>
          <h1>StoreName</h1>
        </Grid>
        <Grid item md={4} container justifyContent="flex-end" style={{color: "grey"}}>
          <Button startIcon={<CircleIcon />} variant="contained" color="error" onClick={GoLive} size="large">Go Live</Button>
        </Grid>
      </Grid>
    );
  }

  return (
    <div>
      <div>
        <Header/>
      </div> 

      <StreamCreatedOverlay />

      <LiveInfoBar />
      <Grid container direction='column'>
          <Grid item container>
            <Grid item xs={false} sm={2} />
            <Grid item xs={12} sm={8}>
                <Productcontent/>
            </Grid>
            <Grid item xs={false} sm={2} />
          </Grid>
      </Grid>  

      <div>
        <Footer/>
      </div>
  </div>
  );
}

export default ProductList;

  
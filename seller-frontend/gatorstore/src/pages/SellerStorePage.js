import React, {useState, useEffect} from 'react';
import Header from '../components/Header';
import Footer from '../components/Footer';
import Button from '@mui/material/Button';
import {Grid} from "@material-ui/core";
import TextField from "@mui/material/TextField";
import List from "@mui/material/List";
import ListItem from "@mui/material/ListItem";
import Productcontent from '../components/ProductContent';

import CircleIcon from '@mui/icons-material/Circle';
import ContentCopyIcon from '@mui/icons-material/ContentCopy';
import '../styles/sellerStorePage.css';

import testStreamObject from '../test-data/streamObject.json';
import sampleProducts from '../test-data/sampleProducts.json';
import gatorPlush from '../images/gator-plush.png';

function SellerStorePage() {
  var storeObject = {id: "2", isLive: false}; // TEST: it's a local replica of the test streamObject.json (represents Database)

  const [liveInfoBarState, SetLiveInfoBarState] = useState('not-live');

  // Checks the status of this store's object in the store API:
  function CheckStoreObject() { 
    /* call API - TODO: When Yiming finishes this API. Store the streamObject in variable storeObject
    const requestOptions = {
      method: 'POST',
      headers: {'jwtToken': jwtToken},
      body: JSON.stringify({ storeID : storeObject.id }) 
    };
    fetch('http://10.136.228.201:8080/store/'+ storeObject.id +'/livestream', requestOptions)
        .then(response => response.json())
        .then(response => {
          if (response.status === 0) {
            SetStreamObject({key: response.result.streamKey, url: response.result.streamUrl});
          } else {
            alert("ERROR: YouTube API did not respond with 'success' status code 0.");
          }
        })
        .catch((error) => {
            console.error(error);
        });*/

    storeObject.isLive = testStreamObject.isLive;
    
    if (storeObject.isLive === true) {
      SetLiveInfoBarState('live');
    } else {
      SetLiveInfoBarState('not-live');
    }
  }

  // Run check-object-function every 60 seconds:
  useEffect(() => {
    const interval = setInterval(() => {
      CheckStoreObject();
    }, 15000);
  
    return () => clearInterval(interval);
  }, [])

  // Hook for overlay
  const [currentOverlay, ChangeCurrentOverlay] = useState("none");
  // Effect for overlay (to freeze scrolling when an overlay is open)
  useEffect(() => {
    document.body.style.overflowY = (currentOverlay !== "none") ? "hidden" : "scroll";    
  }, [currentOverlay]);

  function Overlay() {
    const [streamTitle, SetStreamTitle] = useState("");

    return(
      <div>
        {currentOverlay === 'setStreamTitle' && (
          <div style={{display: 'none'}} />
        )}

        {currentOverlay === 'setStreamTitle' && (
          <div class="go-live-overlay">
            <div class="transparentBG"/>
            <div class="stream-link-container" style={{top: '35vh'}}>
              <p>Enter a title for your stream</p>
              <TextField id="titleField" variant="outlined" color="primary" placeholder="Title" size="small" onChange={e => {SetStreamTitle(e.target.value);}} style={{width: "100%", marginBottom: 15}}/>
              <div style={{textAlign: "center"}}>
                <Button variant="contained" color="warning" onClick={() => {
                    ChangeCurrentOverlay("none");
                }} size="large" style={{marginRight: "10%"}}>Cancel</Button>
                
                <Button variant="contained" color="primary" onClick={() => {
                  if (streamTitle !== "") { // if a title was typed
                      ChangeCurrentOverlay("selectStreamProducts");
                  } else { // give warning that title is required
                    document.getElementById("titleField").style.border = "red solid 2px";
                    document.getElementById("titleField").placeholder = "Title is Required";
                  }
                }} size="large">Continue</Button>
              </div>
            </div>
          </div>
        )}

        {currentOverlay === 'selectStreamProducts' && (
          <div class="go-live-overlay" >
            <div class="transparentBG"/>
            <div class="stream-link-container">
              <h1>What Products Will You Showcase?</h1>
              
              <List selectable={true} selected={0} class="selectStreamItemList">
                {
                  sampleProducts.map(function (product) {
                    return (
                      <ListItem justify="between" class="selectStreamItem" onClick={(e) => {e.target.style.border = "navy solid 2px"}}>
                        <h3>{product.name}</h3>
                        <img src={gatorPlush} />
                        <p>${product.price}</p>
                      </ListItem>
                    );
                  })
                }
              </List>
    
              <div style={{textAlign: "center"}}>
                <Button variant="contained" color="primary" onClick={() => { 
                  // Call YouTube API with this title: 
                  GoLive("My First Test Livestream - Yiming's UF Store");
                  
                  ChangeCurrentOverlay("showStreamCreated");
                }} size="large">Continue</Button>
              </div>
            </div>
          </div>
        )}

        {currentOverlay === 'showStreamCreated' && (
          <div class="go-live-overlay" >
            <div class="transparentBG"/>
            <div class="stream-link-container">
              <h1>Stream Created!</h1>
              <p>URL</p>
              <div class="stream-url-box">
                <p>{newStream.url}</p>
                <Button variant="contained" color="secondary" onClick={() => {navigator.clipboard.writeText(newStream.url)}}><ContentCopyIcon/></Button>
              </div>
    
              <p>Key</p>
              <div class="stream-url-box">
                <p>{newStream.key}</p>
                <Button variant="contained" color="secondary" onClick={() => {navigator.clipboard.writeText(newStream.key)}}><ContentCopyIcon/></Button>
              </div>
    
              <div style={{textAlign: "center"}}>
                <Button variant="contained" color="primary" onClick={() => { ChangeCurrentOverlay("none") }} size="large">Continue</Button>
              </div>
            </div>
          </div>
        )}
      </div>
    )
  }

  // useState is needed for a new stream's information to update that HTML on <StreamCreatedOverlay/> upon their change in GoLive()
  const [newStream, SetStreamObject] = useState({key: "", url: ""});

  const GoLive = async (sTitle) => {
    var jwtToken = window.sessionStorage.getItem("user-jwtToken");
   
   //alert(jwtToken);
    //test:
    //SetStreamObject({key: "32432", url: newStream.url});

    // call API
    const requestOptions = {
      method: 'POST',
      headers: {'Authorization': jwtToken, 'Content-Type': 'application/json'}, // ALWAYS have 'Authorization' with jwtToken in every API
      body: JSON.stringify({ title: sTitle }) 
    };
    fetch('http://10.136.160.70:8080/api/store/'+ storeObject.id +'/livestream', requestOptions)
        .then(response => response.json())
        .then(response => {
          if (response.status === 0) {
            SetStreamObject({key: response.result.streamKey, url: response.result.streamUrl});
          } else {
            alert("ERROR: YouTube API did not respond with 'success' status code 0.");
          }
        })
        .catch((error) => {
            console.error(error);
        });

    // FOR testing purposes ONLY:
    testStreamObject.isLive = true; // this is equivalent to MAKING the Stream Database's 'isLive' field for this store be set to: TRUE
                  // Which is done via the above 'fetch'. Then the function below is called and it checks that 'isLive' field in the Database

    CheckStoreObject(); // to not wait 15 seconds
  }

  function LiveInfoBar() {
    return(
      <div>
        {liveInfoBarState === 'not-live' && (
          <Grid container spacing={0} justifyContent="center" alignItems="center" direction='row' style={{marginBottom: 20}}>
            <Grid item md={4} container>
              <h1>StoreName</h1>
            </Grid>

            <Grid item md={4} container justifyContent="flex-end" style={{color: "grey"}}>
              <Button startIcon={<CircleIcon />} variant="contained" color="error" onClick={() => {
                ChangeCurrentOverlay("setStreamTitle");
              }} size="large">Go Live</Button>
            </Grid>
          </Grid> 
        )}

        {liveInfoBarState === 'live' && (
          <div>
            <Grid container spacing={0} justifyContent="center" alignItems="center" direction='row' style={{marginBottom: 20}}>
              <Grid item md={4} container justifyContent="flex-start">
                <h1>StoreName</h1>
              </Grid>
              <Grid item md={4} container justifyContent="flex-end" style={{color: "red"}}>
                <CircleIcon style={{verticalAlign: 'middle', marginRight: 10}}/>
                <p style={{alignSelf: 'center'}} onClick={() => {testStreamObject.isLive = false;}}><b>LIVE | </b> 3 viewers</p>
              </Grid>
            </Grid>
            <Grid container spacing={0} justifyContent="center" alignItems="center" direction='row' style={{marginBottom: 20}}>
              <Grid item md={4} container justifyContent='flex-start'>
                <iframe width="560" height="315" src="https://www.youtube.com/embed/5qap5aO4i9A?autoplay=1" title="YouTube video player" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe>
              </Grid>
              <Grid item md={4} container>
                <div class="streamChat">
                  <p><b>buyer20:</b> Nice products, mate</p>
                  <p><b>Anon29239824:</b> How much for that one?</p>
                  <p><b>GatorFan1:</b> Noice</p>
                </div>
              </Grid>
            </Grid>
          </div>
        )}
      </div>
    );
  }

  return (
    <div>
      <div>
        <Header/>
      </div> 

      <Overlay/>

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

export default SellerStorePage;

  
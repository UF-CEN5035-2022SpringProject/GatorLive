import React, {useState, useEffect} from 'react';
import Header from '../components/Header';
import Footer from '../components/Footer';
import Button from '@mui/material/Button';
import {Grid} from "@material-ui/core";
import TextField from "@mui/material/TextField";
import Productcontent from '../components/ProductContent';

import CircleIcon from '@mui/icons-material/Circle';
import ContentCopyIcon from '@mui/icons-material/ContentCopy';
import '../styles/sellerStorePage.css';

function ProductList() {
  var storeObject = {id: "2", isLive: false}; // TEST ID
  CheckStoreObject(); // initial check

  const [liveInfoBarState, SetLiveInfoBarState] = useState('not-live');

  // Checks the status of this store's object in the store API:
  function CheckStoreObject() { 
    // storeObject = fetch and stuff from API
    
    if (storeObject.isLive === true) {
      alert("is live!");
      SetLiveInfoBarState('live');
    }
  }

  // Run check-object-function every 60 seconds:
  useEffect(() => {
    const interval = setInterval(() => {
      CheckStoreObject();
    }, 60000);
  
    return () => clearInterval(interval); // This represents the unmount function, in which you need to clear your interval to prevent memory leaks.
  }, [])

  // Overlays' hooks:
  const [successOverlayDisplayed, ChangeSuccessOverlayDisplayed] = useState("none");
  const [titleOverlayDisplayed, ChangeTitleOverlayDisplayed] = useState("none");

  useEffect(() => {
    document.body.style.overflowY = (titleOverlayDisplayed !== "none") ? "hidden" : "scroll";    
  }, [titleOverlayDisplayed]);
  useEffect(() => {
    document.body.style.overflowY = (successOverlayDisplayed !== "none") ? "hidden" : "scroll";
  }, [successOverlayDisplayed]);

  // Overlay displaying input for stream title:
  function StreamTitleOverlay() {
    const [streamTitle, SetStreamTitle] = useState("");

    return(
      <div class="go-live-overlay" style={{display: titleOverlayDisplayed}}>
        <div class="transparentBG"/>
        <div class="stream-link-container" style={{top: '35vh'}}>
          <p>Enter a title for your stream</p>
          <TextField id="titleField" variant="outlined" color="primary" placeholder="Title" size="small" onChange={(e) => SetStreamTitle(e.target.value)} style={{width: "100%", marginBottom: 15}}/>
          <div style={{textAlign: "center"}}>
            <Button variant="contained" color="warning" onClick={() => {
                ChangeTitleOverlayDisplayed("none");
            }} size="large" style={{marginRight: "10%"}}>Cancel</Button>
            
            <Button variant="contained" color="primary" onClick={() => {
              if (streamTitle !== "") { // if a title was typed
                ChangeTitleOverlayDisplayed("none");

                // call YT API with this title: 
                GoLive(streamTitle);

                ChangeSuccessOverlayDisplayed("block"); // display success overlay
              } else { // give warning that title is required
                document.getElementById("titleField").style.border = "red solid 2px";
                document.getElementById("titleField").placeholder = "Title is Required";
              }
            }} size="large">Continue</Button>
          </div>
        </div>
      </div>
    );
  }

  // Overlay displaying url and key for the new stream:
  function StreamCreatedOverlay() {
    return(
      <div class="go-live-overlay" style={{display: successOverlayDisplayed}}>
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
            <Button variant="contained" color="primary" onClick={() => {ChangeSuccessOverlayDisplayed("none") }} size="large">Continue</Button>
          </div>
        </div>
      </div>
    );
  }

  // todo
  // finish fetch for CHECKING IF LIVE (if so, display embed iframe)

  // useState is needed for a new stream's information to update that HTML on <StreamCreatedOverlay/> upon their change in GoLive()
  const [newStream, SetStreamObject] = useState({key: "", url: ""});

  const GoLive = async (sTitle) => {
    var jwtToken = window.sessionStorage.getItem("user-jwtToken");

    //test:
    SetStreamObject({key: "32432", url: sTitle});

    /* call API
    const requestOptions = {
      method: 'POST',
      header: {'jwtToken': jwtToken},
      body: JSON.stringify({ title: sTitle }) 
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

    // FOR testing purposes:
    storeObject.isLive = true;
    CheckStoreObject(); // to not wait 60 seconds maybe? althought the user will prob need a minute to set up the stream on OBS studio
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
                ChangeTitleOverlayDisplayed("block");
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
                <p style={{alignSelf: 'center'}}><b>LIVE | </b> 3 viewers</p>
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

      <StreamTitleOverlay />
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

  
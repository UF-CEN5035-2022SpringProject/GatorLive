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

import settings from '../settings'

function SellerStorePage() {
  var storeObject = {id: "2", isLive: false}; // TEST: it's a local replica of the test streamObject.json (represents Database)

  const [liveInfoBarState, SetLiveInfoBarState] = useState('not-live');
  
  // DONE ONLY ON LOAD: initial check to check if its live or not:
  useEffect(() => {
    CheckStoreObject(); 
  }, []);

  // Checks the status of this store's object in the store API:
  function CheckStoreObject() {
    var jwtToken = window.sessionStorage.getItem("user-jwtToken");

    // call API - TODO: When Yiming finishes this API. Store the embedHTML in variable storeObject
    const requestOptions = {
      method: 'GET',
      headers: {'Authorization': jwtToken}
    };
    fetch(settings.apiHostURL + 'store/' + storeObject.id +'/livestreamStatus', requestOptions)
        .then(response => response.json())
        .then(response => {
          if (response.status === 0) {
            // TODO: SetEmbedHTML(response.result.embedHTML); done when this actually is returned by store status API 
            if (response.result.isLive === true) {
              SetLiveInfoBarState('live');
            } else {
              SetLiveInfoBarState('not-live');
            }
          } else {
            alert("ERROR: YouTube API did not respond with 'success' status code 0.");
          }
        })
        .catch((error) => {
            console.error(error);
        });
  }

  // Hook for overlay
  const [currentOverlay, ChangeCurrentOverlay] = useState("none");
  // Effect for overlay (to freeze scrolling when an overlay is open)
  useEffect(() => {
    document.body.style.overflowY = (currentOverlay !== "none") ? "hidden" : "scroll";    
  }, [currentOverlay]);

  // A twin hook state that holds the same information as streamTitle - it is needed to prevent title reset when switching overlays
  const [newTitle, SetNewTitle] = useState("");

  function Overlay() {
    const [streamTitle, SetStreamTitle] = useState(""); // stream title state

    const [productsSelected, setProductSelected] = React.useState([]); // array for products selected
    
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
                    SetNewTitle(streamTitle);
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
              <h2>Select products to showcase</h2>
              
              <List selected={0} class="selectStreamItemList">
                {
                  sampleProducts.map(function (product) {
                    return (
                      <ListItem selected="false" justify="between" class="selectStreamItem" onClick={
                        (e) => {
                          if (e.currentTarget.style.backgroundColor === "pink") { // being de-selected
                            e.currentTarget.style.boxShadow = "none";
                            e.currentTarget.style.backgroundColor = "lightblue";

                            setProductSelected(productsSelected.filter(item => item !== product.name));
                          } else { // being selected
                            e.currentTarget.style.boxShadow = "inset 0px 0px 0px 2px navy";
                            e.currentTarget.style.backgroundColor = "pink";

                            setProductSelected([...productsSelected, product.name]);
                          }
                        }
                        // Note: e.currentTarget manipulates parent's style (ListItem). e.target manipulates children element's style only.
                      }>
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
                  GetLivestreamKey(newTitle, productsSelected);
                  
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
                <Button variant="contained" color="primary" onClick={() => { 
                  ChangeCurrentOverlay("none");
                  SetLiveInfoBarState('live');
                }} size="large">Go Live</Button>
              </div>
            </div>
          </div>
        )}
      </div>
    )
  }

  // useState is needed for a new stream's information to update that HTML on <StreamCreatedOverlay/> upon their change in GetLivestreamKey()
  const [newStream, SetStreamObject] = useState({key: "", url: ""});

  const GetLivestreamKey = async (sTitle, productList) => {
    var jwtToken = window.sessionStorage.getItem("user-jwtToken");
   
    //alert(productList);
    //test:
    //SetStreamObject({key: "test #", url: sTitle});

    // call API
    const requestOptions = {
      method: 'POST',
      headers: {
        'Authorization': jwtToken
      }, // ALWAYS have 'Authorization' with jwtToken in every API
      body: JSON.stringify({ title: sTitle }) 
    };
    fetch(settings.apiHostURL + 'store/'+ storeObject.id +'/livestream', requestOptions)
      .then(response => response.json())
      .then(response => {
        if (response.status === 0) {
          SetStreamObject({key: response.result.streamKey, url: response.result.streamUrl});
          SetEmbedHTML(response.result.embedHTML);
        } else {
          alert("ERROR: YouTube API did not respond with 'success' status code.");
          window.location.href = "http://localhost:3000/landingpage";
        }
      })
      .catch((error) => {
        console.error(error);
        //alert("ERROR: Back-end is not online or did not respond.");
      });

    // FOR testing purposes ONLY:
    //testStreamObject.isLive = true; // this is equivalent to MAKING the Stream Database's 'isLive' field for this store be set to: TRUE
                  // Which is done via the above 'fetch'. Then the function below is called and it checks that 'isLive' field in the Database

    //CheckStoreObject(); // to not wait 15 seconds
  }

  const [embedHTML, SetEmbedHTML] = useState('');

  function EndLivestream() {
    SetLiveInfoBarState("not-live");

    // some fetch API request telling them that it is no longer live.
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
              }} size="large">Start Livestream</Button>
            </Grid>
          </Grid> 
        )}

        {liveInfoBarState === 'live' && (
          <div>
            <Grid container spacing={0} justifyContent="center" alignItems="center" direction='row' style={{marginBottom: 20}}>
              <Grid item md={4} container justifyContent="flex-start">
                <h1>StoreName</h1>
              </Grid>
              <Grid item md={2} container style={{color: "red", paddingLeft: 10}}>
                <CircleIcon style={{verticalAlign: 'middle', marginRight: 10}}/>
                <p style={{alignSelf: 'center'}} onClick={() => {testStreamObject.isLive = false;}}><b>LIVE | </b> 3 viewers</p>
                
              </Grid>
              <Grid item md={2} container style={{justifyContent: "flex-end"}}>
                <Button variant="contained" color="warning" onClick={() => {
                    EndLivestream();
                  }} size="large">End Livestream</Button>
              </Grid>
              
            </Grid>
            <Grid container spacing={0} justifyContent="center" alignItems="center" direction='row' style={{marginBottom: 20}}>
              <Grid item md={4} container justifyContent='flex-start'>
                
                <div style={{width: 560}} dangerouslySetInnerHTML={{ __html: embedHTML }} />
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

  
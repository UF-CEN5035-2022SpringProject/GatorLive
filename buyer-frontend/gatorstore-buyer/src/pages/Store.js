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
import '../styles/storePage.css';

import testStreamObject from '../test-data/streamObject.json';
import sampleProducts from '../test-data/sampleProducts.json';
import gatorPlush from '../images/gator-plush.png';

import settings from '../settings'

function SellerStorePage() {
  var storeObject = {id: "2", isLive: false}; // TEST: it's a local replica of the test streamObject.json (represents Database)

  const [liveInfoBarState, SetLiveInfoBarState] = useState('not-live');
  
  // On load: initial check to check if its live or not:
  useEffect(() => {
    CheckStoreObject(); 
  }, []);

  // Checks the status of this store's object in the store API:
  function CheckStoreObject() {
    SetLiveInfoBarState('not-live');
    var lofiHTML = '<iframe width="560" height="315" src="https://www.youtube.com/embed/5qap5aO4i9A" title="YouTube video player" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe>';
    lofiHTML = lofiHTML.replace('="560"', '="490"'); // reduce the size of iframe slightly
    SetEmbedHTML(lofiHTML);
    /*var jwtToken = window.sessionStorage.getItem("user-jwtToken");

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
        });*/
  }

  const [embedHTML, SetEmbedHTML] = useState('');

  function LiveInfoBar() {
    return(
      <div>
        {liveInfoBarState === 'not-live' && (
          <Grid container spacing={0} justifyContent="center" alignItems="center" direction='row' style={{marginBottom: 20}}>
            <Grid item md={4} container>
              <h1>The Yiming Store</h1>
            </Grid>
            <Grid item md={4} justifyContent="flex-end" style={{color: "grey"}} container>
                ...is taking a break
            </Grid>
          </Grid> 
        )}

        {liveInfoBarState === 'live' && (
          <div>
            <Grid container spacing={0} justifyContent="center" alignItems="center" direction='row' style={{marginBottom: 20}}>
              <Grid item md={4} container justifyContent="flex-start">
                <h1>The Yiming Store</h1>
              </Grid>
              <Grid item md={4} justifyContent="flex-end" container style={{color: "red", paddingLeft: 10}}>
                <CircleIcon style={{verticalAlign: 'middle', marginRight: 10}}/>
                <p style={{alignSelf: 'center'}} onClick={() => {testStreamObject.isLive = false;}}><b>LIVE | </b> 3 viewers</p>
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

  
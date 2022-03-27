import React, {useState, useEffect} from 'react';
import Header from '../components/Header';
import Footer from '../components/Footer';
import Button from '@mui/material/Button';
import {Grid} from "@material-ui/core";
import TextField from "@mui/material/TextField";
import List from "@mui/material/List";
import ListItem from "@mui/material/ListItem";
import Productcontent from '../components/ProductContent';
import Productcard from '../components/Productcard';

import CircleIcon from '@mui/icons-material/Circle';
import '../styles/storePage.css';

import testStreamObject from '../test-data/streamObject.json';
import sampleProducts from '../test-data/sampleProducts.json';
import gatorPlush from '../images/gator-plush.png';
import settings from '../settings'

import { useParams } from 'react-router-dom';

function SellerStorePage() {
  var storeObject = {id: "gatorstore-1", isLive: false}; // TEST: it's a local replica of the test streamObject.json (represents Database)

  const [liveInfoBarState, SetLiveInfoBarState] = useState('not-live');

  const { storeID } = useParams(); // Get StoreID string from the url
  
  // On load: initial check to check if its live or not:
  useEffect(() => {
    CheckStoreObject(); 
  }, []);

  // Every 15 seconds:
  useEffect(() => {
    const interval = setInterval(() => {
      console.log("5-second Fetch...")
      CheckStoreObject(); 
    }, 5000);
  
    return () => clearInterval(interval);
  }, [])

  // Checks the status of this store's object in the store API:
  function CheckStoreObject() {
    /* TESTING:
    SetLiveInfoBarState('live');
    var lofiHTML = '<iframe width="560" height="315" src="https://www.youtube.com/embed/5qap5aO4i9A" title="YouTube video player" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe>';
    lofiHTML = lofiHTML.replace('="560"', '="490"'); // reduce the size of iframe slightly
    SetEmbedHTML(lofiHTML);

    var lofiChatHTML = '<iframe width="494" height="315" src="https://www.youtube.com/live_chat?v=5qap5aO4i9A&embed_domain=localhost" frameborder="0"></iframe>';
    SetEmbedChatHTML(lofiChatHTML);*/

    var jwtToken = window.sessionStorage.getItem("user-jwtToken");

    // call API - TODO: When Yiming finishes this API. Store the embedHTML in variable storeObject
    const requestOptions = {
      method: 'GET',
      headers: {'Authorization': jwtToken}
    };
    fetch(settings.apiHostURL + 'store/' + storeObject.id, requestOptions)
        .then(response => response.json())
        .then(response => {
          if (response.status === 0) {
            // TODO: SetEmbedHTML(response.result.embedHTML); done when this actually is returned by store status API 
            if (response.result.isLive === true) {
              SetLiveInfoBarState('live');

              var embedStreamHTML = '<iframe width="490" height="315" src="https://www.youtube.com/embed/' + response.result.liveId + '" title="YouTube video player" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe>'
              SetEmbedHTML(embedStreamHTML);
              var embedChatRoomHTML = '<iframe width="494" height="315" src="https://www.youtube.com/live_chat?v=' + response.result.liveId + '&embed_domain=localhost" frameborder="0"></iframe>';
              SetEmbedChatHTML(embedChatRoomHTML);
              // Note: The embed HTML for the chat specifies the host in which to run it in. Currently "localhost"!
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

  const [embedHTML, SetEmbedHTML] = useState('');
  const [embedChatHTML, SetEmbedChatHTML] = useState('');

  function LiveInfoBar() {
    return(
      <div>
        {liveInfoBarState === 'not-live' && (
          <Grid container spacing={0} justifyContent="center" alignItems="center" direction='row' style={{marginBottom: 20}}>
            <Grid item md={4} container>
              <h1>{storeID}</h1>
            </Grid>
            <Grid item md={4} justifyContent="flex-end" style={{color: "grey"}} container>
                ...is taking a break
            </Grid>
          </Grid> 
        )}

        {liveInfoBarState === 'live' && (
          <div>
            <Grid container spacing={0} justifyContent="center" alignItems="center" direction='row' style={{marginBottom: 15}}>
              <Grid item md={4} container justifyContent="flex-start">
                <h1>The Yiming Store</h1>
              </Grid>
              <Grid item md={4} justifyContent="flex-end" container style={{color: "red", paddingLeft: 10}}>
                <CircleIcon style={{verticalAlign: 'middle', marginRight: 10}}/>
                <p style={{alignSelf: 'center'}} onClick={() => {testStreamObject.isLive = false;}}><b>LIVE | </b> 3 viewers</p>
              </Grid>
            </Grid>

            <Grid container spacing={0} justifyContent="center" alignItems="center" direction='row' style={{marginBottom: 0}}>
              <Grid item md={4} container justifyContent='flex-start'>
                <div dangerouslySetInnerHTML={{ __html: embedHTML }} />
              </Grid>
              <Grid item md={4} container justifyContent='flex-start'>
                <div dangerouslySetInnerHTML={{ __html: embedChatHTML }} />
              </Grid>
            </Grid>

            <Grid container spacing={0} justifyContent="center" alignItems="center" direction='row' style={{marginBottom: 20}}>
              <Grid item md={8} container direction='column' justifyContent='flex-start' style={{backgroundColor: "#202020", padding: "10px 15px 0px"}}>
                <div class="featuredItemTitle">Featured Items</div>
                <List selected={0} class="selectStreamItemList">
                {sampleProducts.map(function (product) {
                  return (
                      <ListItem selected="false" justify="between" class="selectStreamItem" style={{backgroundColor: "rgb(226, 197, 164)"}} onClick={
                        (e) => {
                          
                        }
                        // Note: e.currentTarget manipulates parent's style (ListItem). e.target manipulates children element's style only.
                      }>
                        <div>{product.name}</div>
                        <img src={gatorPlush} />
                        <p>${product.price}</p>
                      </ListItem>
                      );
                    })
                  }
                </List>
              </Grid>
            </Grid>
          </div>
        )}
      </div>
    );
  }

  const [productArray, SetProductArray] = useState([]);

  const [currProductPage, ChangeProductPage] = useState(0);
  var maxProductPage = 1; // default

  function ProductList() {
    // Calls on GetPage() to get a new product page upon the user scrolling down.
    function ScrollDown() {
      // Only request more products if current page number is below max:
      if (currProductPage <= maxProductPage) {
        ChangeProductPage(currProductPage + 1);
        GetPage(currProductPage);
      }
    }

    function GetPage(pageNum) {
      // Get JWT Token for POST request header:
      var jwtToken = window.sessionStorage.getItem("user-jwtToken");

      TODO:
      /* Call API to get product list:
      const requestOptions = {
        method: 'GET',
        headers: {
          'Authorization': jwtToken
        },
        body: {}
      };
      fetch(settings.apiHostURL + 'store/' + storeID + '/productList?page=' + pageNum, requestOptions)
        .then(response => response.json())
        .then(response => {
          if (response.status === 0) {
            // if page requested isn't more than max page: Add products of this new page to "productArray"
            TODO
            // Set max page number so that we know it from the first load:
            maxProductPage = response.result.maxPage;
          }
          else {
            alert("ERROR: Product Page API did not respond with 'success' status code.");
            window.location.href = "http://localhost:3001/";
          }
        })
        .catch((error) => {
          console.error(error);
        });*/
    }
    
    return(
      <div>
        <Grid container spacing={2}>
        {productArray.map(function (product) {
            return(
              <Grid item xs={12} sm={4}>
                <Productcard 
                  title="Title" 
                  subtitle="Price" 
                  imageUrl="https://media.wired.com/photos/5f23168c558da0380aa8e37f/master/pass/Gear-Google-Pixel-4A-front-and-back-angle-SOURCE-Google.jpg"
                  description="Unlocked Android phone gives you the flexibility to change carriers and choose your own data plan; works with Verizon, T-Mobile, Sprint, AT&T, Google Fi, and other major carriers"
                />
              </Grid>
            );
          })
        }
        </Grid>
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

  
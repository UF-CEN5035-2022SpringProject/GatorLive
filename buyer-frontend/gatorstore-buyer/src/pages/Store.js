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
  const [liveInfoBarState, SetLiveInfoBarState] = useState('not-live');

  const [storeName, SetStoreName] = useState("Store!"); // to replace after the first fetch of the store object

  const { storeID } = useParams(); // Get StoreID string from the url
  
  // On load: initial check to check if its live or not:
  useEffect(() => {
    CheckStoreObject(true); 
  }, []);

  // Every 15 seconds:
  useEffect(() => {
    const interval = setInterval(() => {
      console.log("5-second Fetch...")
      CheckStoreObject(false); 
    }, 5000);
  
    return () => clearInterval(interval);
  }, [])

  // Live Product List Hook array:
  const [liveProductArray, SetLiveProductArray] = useState([{
    name: "Test Live Product",
    price: "$3.50",
    description: "Buy it!"
  }]);

  // Checks the status of this store's object in the store API:
  function CheckStoreObject(detail) {
    /* TESTING:
    SetLiveInfoBarState('live');
    var lofiHTML = '<iframe width="560" height="315" src="https://www.youtube.com/embed/5qap5aO4i9A" title="YouTube video player" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe>';
    lofiHTML = lofiHTML.replace('="560"', '="490"'); // reduce the size of iframe slightly
    SetEmbedHTML(lofiHTML);

    var lofiChatHTML = '<iframe width="494" height="315" src="https://www.youtube.com/live_chat?v=5qap5aO4i9A&embed_domain=localhost" frameborder="0"></iframe>';
    SetEmbedChatHTML(lofiChatHTML);*/

    var jwtToken = window.sessionStorage.getItem("user-jwtToken");

    // call API:
    const requestOptions = {
      method: 'GET',
      headers: {'Authorization': jwtToken}
    };
    fetch(settings.apiHostURL + 'store/' + storeID + '/livestreamStatus?detail=' + detail, requestOptions)
        .then(response => response.json())
        .then(response => {
          if (response.status === 0) {
            SetStoreName(response.result.name); // get name of store

            if (response.result.isLive === true) {
              SetLiveInfoBarState('live');

              var embedStreamHTML = '<iframe width="490" height="315" src="https://www.youtube.com/embed/' + response.result.liveId + '" title="YouTube video player" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe>'
              SetEmbedHTML(embedStreamHTML);
              var embedChatRoomHTML = '<iframe width="494" height="315" src="https://www.youtube.com/live_chat?v=' + response.result.liveId + '&embed_domain=localhost" frameborder="0"></iframe>';
              SetEmbedChatHTML(embedChatRoomHTML);
              // Note: The embed HTML for the chat specifies the host in which to run it in. Currently "localhost"!

              if (detail === true) {
                // Add live products to some hook array:
                SetLiveProductArray(response.result.productList);
              }
            } else {
              SetLiveInfoBarState('not-live');
            }
          } else {
            console("ERROR: Store Object API did not respond with 'success' status code 0.");
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
              <h1>{storeName}</h1>
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
                <h1>{storeName}</h1>
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
                {liveProductArray.map(function (product) {
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

  const [productArray, SetProductArray] = useState([
    {
      name: "Product",
      price: "$3.50",
      description: "This is a very nice product to buy for a special occassion. DM me for offers!"
    }
  ]);

  const [currProductPage, ChangeProductPage] = useState(0);
  var maxProductPage = 1; // default

  function ProductList() {
    GetPage(0); // get first page automatically

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
      
      // Call API to get product list:
      const requestOptions = {
        method: 'GET',
        headers: {
          'Authorization': jwtToken
        },
        body: {}
      };
      fetch(settings.apiHostURL + 'store/' + storeID + '/product-list?page=' + pageNum, requestOptions)
        .then(response => response.json())
        .then(response => {
          if (response.status === 0) {
            // if page requested isn't more than max page: Add products of this new page to "productArray"
            if (pageNum <= response.result.maxPage) {
              SetProductArray(productArray.concat(response.result.productList));
            }

            // Set max page number so that this fetch isn't even called if it is an invalid page number
            maxProductPage = response.result.maxPage;
          }
          else {
            alert("ERROR: Product Page API did not respond with 'success' status code.");
            window.location.href = "http://localhost:3001/";
          }
        })
        .catch((error) => {
          console.error(error);
        });
    }
    
    return(
      <div class="product-container" onScroll={(e) => {
        if (e.target.scrollHeight - e.target.scrollTop === e.target.clientHeight) {
          ScrollDown();
        }
      }}>
        <Grid container spacing={2}>
          {productArray && productArray.length > 0 && ( productArray.map(function (product) {
                return(
                  <Grid item xs={12} sm={4}>
                    <Productcard 
                      title= {product.name} 
                      subtitle={product.price}
                      imageUrl="https://media.wired.com/photos/5f23168c558da0380aa8e37f/master/pass/Gear-Google-Pixel-4A-front-and-back-angle-SOURCE-Google.jpg"
                      description={product.description}
                    />
                  </Grid>
                );
              })
          )}
          {productArray.length == 0 && (
            <div>- No products here -</div>
          )}
        </Grid>
      </div>
    );
  }

  return (
    <div className="RootFlexContainer">
      <div>
        <Header/>
      </div> 

      <div style={{minHeight: "80vh"}}>
        <LiveInfoBar />
        
        <Grid container direction='column'>
          <Grid item container>
            <Grid item xs={false} sm={2} />
            <Grid item xs={12} sm={8}>
              <ProductList/>
            </Grid>
            <Grid item xs={false} sm={2} />
          </Grid>
        </Grid>  
      </div>

      <div>
        <Footer/>
      </div>
  </div>
  );
}

export default SellerStorePage;

  
import React, {useState, useEffect} from 'react';
import Header from '../components/Header';
import Footer from '../components/Footer';
import Button from '@mui/material/Button';
import {Grid} from "@material-ui/core";
import TextField from "@mui/material/TextField";
import List from "@mui/material/List";
import ListItem from "@mui/material/ListItem";
import Productcard from '../components/Productcard';

import CircleIcon from '@mui/icons-material/Circle';
import AddIcon from '@mui/icons-material/Add';
import ContentCopyIcon from '@mui/icons-material/ContentCopy';
import '../styles/sellerStorePage.css';

import testStreamObject from '../test-data/streamObject.json';
import sampleProducts from '../test-data/sampleProducts.json';
import gatorPlush from '../images/gator-plush.png';

import settings from '../settings'
import { useParams } from 'react-router-dom';
import { useNavigate } from "react-router-dom";

function SellerStorePage() {
  const navigate = useNavigate(); // to redirect using navigate()

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
    var jwtToken = window.sessionStorage.getItem("user-jwtToken");

    // call API - TODO: When Yiming finishes this API. Store the embedHTML in variable storeObject
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

              if (detail === true) {
                // Add live products to some hook array:
                SetLiveProductArray(response.result.productList);
              }
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

    // Product Creation Hooks:
    const [prodTitle, SetProdTitle] = useState("");
    const [prodPrice, SetProdPrice] = useState("");
    const [prodDescription, SetProdDescription] = useState("");
    
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

        {currentOverlay === 'createNewProduct' && (
          <div class="go-live-overlay" >
            <div class="transparentBG"/>
            <div class="stream-link-container" style={{top: "18vh"}}>
              <h1>Create New Product:</h1>
              <p>Title</p>
              <TextField id="prodTitleField" variant="outlined" color="primary" placeholder="Product Title" size="small" onChange={e => {SetProdTitle(e.target.value);}} style={{width: "100%", marginBottom: 15}}/>
              <p>Price</p>
              <TextField id="prodPriceField" variant="outlined" color="primary" placeholder="Product Price" size="small" onChange={e => {SetProdPrice(e.target.value);}} style={{width: "100%", marginBottom: 15}}/>
              <p>Description</p>
              <TextField id="prodDescriptionField" variant="outlined" color="primary" placeholder="Product Description" size="small" onChange={e => {SetProdDescription(e.target.value);}} style={{width: "100%", marginBottom: 15}}/>
              <div style={{display: "flex", flexDirection: "row", justifyContent: "center"}}>
                <p style={{padding: "5px 10px 0px"}}>Image: </p>
                <Button variant="contained" color="secondary" onClick={() => {
                  alert("Sir, this is a dummy button for now.")
                }} size="small">Upload File</Button>
              </div>

              <div style={{textAlign: "center", paddingTop: "17px"}}>
                <Button variant="contained" color="warning" onClick={() => {
                    ChangeCurrentOverlay("none");
                }} size="large" style={{marginRight: "10%"}}>Cancel</Button>

                <Button variant="contained" color="primary" onClick={() => {
                  if (prodTitle == "") {
                    document.getElementById("prodTitleField").style.border = "red solid 2px";
                    document.getElementById("prodTitleField").placeholder = "Title is Required";
                  }
                  if (prodPrice == "") {
                    document.getElementById("prodPriceField").style.border = "red solid 2px";
                    document.getElementById("prodPriceField").placeholder = "Product Price is Required";
                  } 
                  if (prodDescription == "") {
                    document.getElementById("prodDescriptionField").style.border = "red solid 2px";
                    document.getElementById("prodDescriptionField").placeholder = "Description is Required";
                  }

                  if (prodTitle != "" && prodPrice != "" && prodDescription != "") {
                    // Call Create Product API with this info:
                    CreateNewProduct(prodTitle, prodPrice, prodDescription, 1, "No Picture");
                  }
                }} size="large">Create Product</Button>
              </div>
            </div>
          </div>
        )}
      </div>
    )
  }

  function CreateNewProduct(prodTitle, prodPrice, prodDescription, prodQuantity, prodPic) {
    var jwtToken = window.sessionStorage.getItem("user-jwtToken");

    const requestOptions = {
      method: 'POST',
      headers: {'Authorization': jwtToken},
      body: JSON.stringify({ name: prodTitle, price: prodPrice, description: prodDescription, quantity: prodQuantity, picture: prodPic, storeId: storeID}) 
    };
    fetch(settings.apiHostURL + 'product/create/', requestOptions)
        .then(response => response.json())
        .then(response => {
          if (response.status === 0) {
            alert("Product created on API. Now redirecting")
            navigate("/product/" + response.result.id);
          } else {
            alert("ERROR: YouTube API did not respond with 'success' status code 0.");
          }
        })
        .catch((error) => {
            console.error(error);
        });
  }

  // useState is needed for a new stream's information to update that HTML on <StreamCreatedOverlay/> upon their change in GetLivestreamKey()
  const [newStream, SetStreamObject] = useState({key: "", url: ""});

  const GetLivestreamKey = async (sTitle, productList) => {
    var jwtToken = window.sessionStorage.getItem("user-jwtToken");

    const requestOptions = {
      method: 'POST',
      headers: {
        'Authorization': jwtToken
      }, 
      body: JSON.stringify({ title: sTitle, productIdList: productList }) 
    };
    fetch(settings.apiHostURL + 'store/'+ storeID +'/livestream', requestOptions)
      .then(response => response.json())
      .then(response => {
        if (response.status === 0) {
          SetStreamObject({key: response.result.streamKey, url: response.result.streamUrl});
          SetEmbedHTML(response.result.embedHTML);
        } else {
          alert("ERROR: YouTube API did not respond with 'success' status code.");
          window.location.href = "http://localhost:3000/";
        }
      })
      .catch((error) => {
        console.error(error);
      });

    // FOR testing purposes ONLY:
    //testStreamObject.isLive = true; // this is equivalent to MAKING the Stream Database's 'isLive' field for this store be set to: TRUE
                  // Which is done via the above 'fetch'. Then the function below is called and it checks that 'isLive' field in the Database
  }

  const [embedHTML, SetEmbedHTML] = useState('');
  const [embedChatHTML, SetEmbedChatHTML] = useState('');

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
              <h1>{storeName}</h1>
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

  function ProductsHeader() {
    return(
      <Grid container spacing={0} justifyContent="center" alignItems="center" direction='row' style={{marginBottom: 20}}>
        <Grid item md={4} container>
          <h2>Products:</h2>
        </Grid>
        <Grid item md={4} justifyContent="flex-end" container>
          <Button size="medium" startIcon={<AddIcon />} variant="contained" color="primary" onClick={() => {
            ChangeCurrentOverlay("createNewProduct");
          }}>Create Product</Button>
        </Grid>
      </Grid>
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

      <Overlay/>

      <div style={{minHeight: "80vh"}}>
        <LiveInfoBar />

        <ProductsHeader/>
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

  
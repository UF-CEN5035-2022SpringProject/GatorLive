import React, {useState, useEffect} from 'react';
import Header from '../components/Header';
import Footer from '../components/Footer';
import '../styles/sellerStoreList.css';

import AddIcon from '@mui/icons-material/Add';
import EditIcon from '@mui/icons-material/Edit';
import DeleteIcon from '@mui/icons-material/Delete';
import OpenInNewIcon from '@mui/icons-material/OpenInNew';
import Button from '@mui/material/Button';
import TextField from "@mui/material/TextField";
import Avatar from '@mui/material/Avatar';

import settings from '../settings'

function StoreEntry(storeInfo) {
  return (
    <div className="StoreEntry flexCenter rowFlex">
      <div style={{ flex: 0.7 }} className="flexCenter">
        <Avatar sx={{ bgcolor: 'navy', width: 80, height: 80}}>{storeInfo.initials}</Avatar>
      </div>
      <div style={{ flex: 3 }}>
        <h1 className="StoreEntryTitle">{storeInfo.name} <OpenInNewIcon /></h1>
        <div className="StoreEntryDetailRow">
          <p>{storeInfo.activeListings} Active Listings</p>
          <p>|</p>
          <p>{storeInfo.unfinishedOrders} Unfinished Orders</p>
        </div>
      </div>
      <div style={{ flex: 1}} className="flexCenter">
        <Button startIcon={<EditIcon />} variant="contained" color="primary" sx={{ marginBottom: 1 }}>Edit</Button>
        <Button startIcon={<DeleteIcon />} variant="contained" color="secondary">Delete</Button>
      </div>
    </div>
  );
}

function SellerStoreList() {
  // Hook for overlay
  const [currentOverlay, ChangeCurrentOverlay] = useState("none");
  // Effect for overlay (to freeze scrolling when an overlay is open)
  useEffect(() => {
    document.body.style.overflowY = (currentOverlay !== "none") ? "hidden" : "auto";    
  }, [currentOverlay]);

  // Twin hook to get the store name out of Overlay():
  const [newStore, SetNewStore] = useState("");

  function Overlay() {
    const [newStoreName, SetNewStoreName] = useState(""); // state for name for new store

    return (
      <div>
        {currentOverlay === 'none' && (
          <div style={{display: 'none'}} />
        )}

        {currentOverlay === 'createNewStore' && (
          <div class="create-store-overlay">
            <div class="transparentBG"/>
            <div class="overlay-container" style={{top: '35vh'}}>
              <p>Enter a name for your new store:</p>
              <TextField id="NameField" variant="outlined" color="primary" placeholder="Store Name" size="small" onChange={e => {SetNewStoreName(e.target.value);}} style={{width: "100%", marginBottom: 15}}/>
              <div style={{textAlign: "center"}}>
                  <Button variant="contained" color="warning" onClick={() => {
                      ChangeCurrentOverlay("none");
                  }} size="large" style={{marginRight: "10%"}}>Cancel</Button>
                  
                  <Button variant="contained" color="primary" onClick={() => {
                    if (newStoreName !== "") { // if a name was typed
                      SetNewStore(newStoreName);

                      // Create store:
                      CreateNewStore(newStoreName);

                      ChangeCurrentOverlay("newStoreConfirmation");
                    } else { // give warning that name is required
                      document.getElementById("NameField").style.border = "red solid 2px";
                      document.getElementById("NameField").placeholder = "A store name is required";
                    }
                  }} size="large">Continue</Button>
                </div>
            
            </div>
          </div>
        )}

        {currentOverlay === 'newStoreConfirmation' && ( 
          <div class="create-store-overlay">
            <div class="transparentBG"/>
            <div class="overlay-container" style={{top: '35vh'}}>
              <h3><b>{newStore}</b> was successfully created!</h3>

              <div style={{textAlign: "center"}}>
                <Button variant="contained" color="warning" onClick={() => {
                    ChangeCurrentOverlay("none");
                }} size="large">Ok</Button>
              </div>
            </div>
          </div>
        )}
      </div>
    );
  }

  function CreateNewStore(name) {
    /* Get JWT Token for POST request header:
    var jwtToken = window.sessionStorage.getItem("user-jwtToken");

    // Call API:
    const requestOptions = {
      method: 'POST',
      headers: {
        'Authorization': jwtToken
      },
      body: JSON.stringify({ name: name }) 
    };
    fetch(settings.apiHostURL + 'store/create', requestOptions)
      .then(response => response.json())
      .then(response => {
        if (response.status !== 0) {
          alert("ERROR: Create Store API did not respond with 'success' status code.");
          window.location.href = "http://localhost:3000/";
        }
      })
      .catch((error) => {
        console.error(error);
        //alert("ERROR: Back-end is not online or did not respond.");
      });*/
  }


  return (
    <div className="RootFlexContainer">
      <Header />

      <Overlay/>

      <div className="flexCenter colFlex">
        <div className="storeListSubHeader">
          <h1>Your Stores</h1>
          <Button startIcon={<AddIcon/>} color="primary" variant="contained" size="medium" onClick={() => {
            ChangeCurrentOverlay("createNewStore");
          }}>New Store</Button>
        </div>
        <div className="storeListContainer">
          <StoreEntry initials="YS" name="Yiming Store" activeListings={3} unfinishedOrders={2}/>
          <StoreEntry initials="UF" name="UF Shop" activeListings={1} unfinishedOrders={4}/>
          <StoreEntry initials="TS" name="Test Store" activeListings={0} unfinishedOrders={1}/>
          
        </div>
      </div>
      <Footer />
    </div>
  );
}

export default SellerStoreList;
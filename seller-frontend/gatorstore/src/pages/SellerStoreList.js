import React from 'react';
import Header from '../components/Header';
import Footer from '../components/Footer';
import '../styles/sellerStoreList.css';

import AddIcon from '@mui/icons-material/Add';
import EditIcon from '@mui/icons-material/Edit';
import DeleteIcon from '@mui/icons-material/Delete';
import OpenInNewIcon from '@mui/icons-material/OpenInNew';
import Button from '@mui/material/Button';
import Avatar from '@mui/material/Avatar';

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

export default function SellerStoreList() {
  return (
    <div className="RootFlexContainer">
      <Header />
      <div className="flexCenter colFlex">
        <div className="storeListSubHeader">
          <h1>Your Stores</h1>
          <Button startIcon={<AddIcon/>} color="primary" variant="contained" size="medium">New Store</Button>
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

import React from 'react';
import Header from '../components/Header';
import Footer from '../components/Footer';
import '../styles/sellerProductPage.css';

import gatorMug from '../images/gator.jpg';
import EditIcon from '@mui/icons-material/Edit';
import MoveUpIcon from '@mui/icons-material/MoveUp';
import Button from '@mui/material/Button';
import Avatar from '@mui/material/Avatar';

export default function ProductPage() {
    return (
        <div className="RootFlexContainer">
            <Header />
            <div className="ProductDetailsRow">
                <div className="ProductImage">
                    <img src={gatorMug} alt="Gator Mug"/>
                    <Button color="primary" variant="contained"><EditIcon /></Button>
                </div>
                <div className="ProductKeyDetailColumn">
                    <h1>Gator Mug</h1>
                    <h2>$5.99</h2>
                    <Button startIcon={<EditIcon />} color="primary" variant="contained" size="large">Edit Details</Button>
                </div>
                <div style={{flex: 1}} className="flexCenter">
                    <div className="ProductStoreInfo flexCenter colFlex">
                        <h1>UF Store</h1>
                        <Avatar sx={{ bgcolor: 'navy', fontSize: 50, width: 120, height: 120}}>UF</Avatar>
                        <Button startIcon={<MoveUpIcon />} color="secondary" variant="contained" size="large" sx={{marginBottom: 3}}>Switch to Another Store</Button>
                        <Button startIcon={<EditIcon />} color="primary" variant="contained" size="large" sx={{marginBottom: 1}}>Edit Store</Button>
                    </div>
                </div>
            </div>
            <div className="ProductDescription">
                This is the description for this product. You should buy this.
            </div>
            <Footer />
        </div>
    );
}
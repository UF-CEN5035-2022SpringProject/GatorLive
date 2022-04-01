import React, {useState, useEffect} from 'react';
import Header from '../components/Header';
import Footer from '../components/Footer';
import '../styles/sellerProductPage.css';

import gatorMug from '../images/gator.jpg';
import EditIcon from '@mui/icons-material/Edit';
import MoveUpIcon from '@mui/icons-material/MoveUp';
import Button from '@mui/material/Button';
import Avatar from '@mui/material/Avatar';

import settings from '../settings'
import { useParams, Link } from 'react-router-dom';

export default function ProductPage() {
    const [productInfo, SetProductInfo] = useState(null);
    // Effect for overlay (to freeze scrolling when an overlay is open)
    useEffect(() => {
        getProductInfo()
    }, [])
    
    const { productID } = useParams(); // Get StoreID string from the url
    function getProductInfo() {
        console.log('Get Product ID:' + productID);
        const requestOptions = {
            method: 'GET'
        };
        fetch(settings.apiHostURL + `product/${productID}/info`, requestOptions) // PA2
        .then(response => response.json())
        .then(response => {
            if (response.status === 0) {
                SetProductInfo(response.result);
            }
            else {
                alert("ERROR: YouTube API did not respond with 'success' status code 0.");
            }
        })
        .catch((error) => {
            console.error(error);
        });
    }
    
    return (
        <div className="RootFlexContainer">
            <Header />
            <div className="ProductDetailsRow">
                <div className="ProductImage">
                    <img src={gatorMug} alt="Gator Mug"/>
                    <Button color="primary" variant="contained"><EditIcon /></Button>
                </div>
                {productInfo ? 
                    <div className="ProductKeyDetailColumn">
                        <h1>{productInfo.name}</h1>
                        <h2>{`${productInfo.price}`}</h2>
                        <Button startIcon={<EditIcon />} color="primary" variant="contained" size="large">Edit Details</Button>
                    </div> 
                    : 
                    null
                }
                
                <div style={{flex: 1}} className="flexCenter">
                    <div className="ProductStoreInfo flexCenter colFlex">
                        <h1>GatorStore</h1>
                        <Avatar sx={{ bgcolor: 'navy', fontSize: 50, width: 120, height: 120}}>UF</Avatar>
                        <Button component={Link} to={'/store-list'} startIcon={<MoveUpIcon />} color="secondary" variant="contained" size="large" sx={{marginBottom: 3}}>Back to Store List</Button>
                        {/* <Button startIcon={<EditIcon />} color="primary" variant="contained" size="large" sx={{marginBottom: 1}}>Edit Store</Button> */}
                    </div>
                </div>
            </div>
            {productInfo?
                <div className="ProductDescription">
                    <h2>Product Description</h2>
                    <br></br>
                    {productInfo.description}
                </div>
                :
                null
            }
            <Footer />
        </div>
    );
}
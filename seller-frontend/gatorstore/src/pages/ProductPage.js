import React, {useState, useEffect} from 'react';
import Header from '../components/Header';
import Footer from '../components/Footer';
import '../styles/sellerProductPage.css';

import gatorMug from '../images/gator.jpg';
import EditIcon from '@mui/icons-material/Edit';
import MoveUpIcon from '@mui/icons-material/MoveUp';
import Button from '@mui/material/Button';
import TextField from "@mui/material/TextField";
import Avatar from '@mui/material/Avatar';
import DeleteIcon from '@mui/icons-material/Delete';

import settings from '../settings'
import '../styles/sellerStorePage.css';
import { useParams, Link } from 'react-router-dom';
import { useNavigate } from "react-router-dom";

export default function ProductPage() {
    const navigate = useNavigate(); // to redirect using navigate()

    const [productInfo, SetProductInfo] = useState(null);
    // On load: get product info:
    useEffect(() => {
        getProductInfo()
    }, [])
    
    const { productID } = useParams(); // Get ProductID string from the url
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

    function UpdateProductInfo(title, price, desc, quantity, picture) {
        var jwtToken = window.sessionStorage.getItem("user-jwtToken");
        const requestOptions = {
            method: 'PUT',
            headers: {'Authorization': jwtToken},
            body: JSON.stringify({name: title, price: price, description: desc, quantity: quantity, picture: picture})
        };
        fetch(settings.apiHostURL + `product/${productID}`, requestOptions)
        .then(response => response.json())
        .then(response => {
            if (response.status === 0) {
                getProductInfo();
                ChangeCurrentOverlay("none");
            }
            else {
                console.log("ERROR: Update product API did not respond with 'success' status code 0.");
            }
        })
        .catch((error) => {
            console.error(error);
        });
    }

    function DeleteProduct() {
        var jwtToken = window.sessionStorage.getItem("user-jwtToken");
        const requestOptions = {
            method: 'DELETE',
            headers: {'Authorization': jwtToken}
        };
        fetch(settings.apiHostURL + `product/${productID}`, requestOptions)
        .then(response => response.json())
        .then(response => {
            if (response.status === 0) {
                navigate("/store/" + productInfo.storeId); // redirect to product's store
            }
            else {
                console.log("ERROR: Delete product API did not respond with 'success' status code 0.");
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

    function Overlay() {
        // Product Creation Hooks:
        const [prodTitle, SetProdTitle] = useState("");
        const [prodPrice, SetProdPrice] = useState("");
        const [prodDescription, SetProdDescription] = useState("");

        return(
            <div>
                {currentOverlay === 'none' && (
                    <div style={{display: 'none'}} />
                )}

                {currentOverlay === 'updateProduct' && (
                    <div class="go-live-overlay" >
                        <div class="transparentBG"/>
                        <div class="stream-link-container" style={{top: "18vh"}}>
                        <h1>Update Product:</h1>
                        <p>Title</p>
                        <TextField id="prodTitleField" placeholder={productInfo.name} variant="outlined" color="primary" size="small" onChange={e => {SetProdTitle(e.target.value);}} style={{width: "100%", marginBottom: 15}}/>
                        <p>Price</p>
                        <TextField id="prodPriceField" placeholder={productInfo.price} type="number" variant="outlined" color="primary" size="small" onChange={e => {SetProdPrice(e.target.value);}} style={{width: "100%", marginBottom: 15}}/>
                        <p>Description</p>
                        <TextField id="prodDescriptionField" placeholder={productInfo.description} variant="outlined" color="primary" size="small" onChange={e => {SetProdDescription(e.target.value);}} style={{width: "100%", marginBottom: 15}}/>
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
                                // Call Update Product API
                                UpdateProductInfo(prodTitle, parseFloat(prodPrice), prodDescription, 1, "No Picture");
                            }
                            }} size="large">Update Product</Button>
                        </div>
                        </div>
                    </div>
                )}

                {currentOverlay === 'deleteConfirm' && (
                    <div class="go-live-overlay" >
                        <div class="transparentBG"/>
                        <div class="stream-link-container" style={{top: "35vh"}}>
                        <h2>Are you sure you want to delete "{productInfo.name}"?</h2>

                        <div style={{textAlign: "center", paddingTop: "17px"}}>
                            <Button variant="contained" color="warning" onClick={() => {
                                ChangeCurrentOverlay("none");
                            }} size="large" style={{marginRight: "10%"}}>Cancel</Button>

                            <Button variant="contained" color="error" onClick={() => {
                                DeleteProduct();
                            }} size="large">Yes</Button>
                        </div>
                        </div>
                    </div>
                )}
            </div>
        );
    }    
    
    return (
        <div className="RootFlexContainer">
            <Header />

            <Overlay/>
            <div className="ProductDetailsRow">
                <div className="ProductImage">
                    <img src={gatorMug} alt="Gator Mug"/>
                    <Button color="primary" variant="contained"><EditIcon /></Button>
                </div>
                {productInfo ? 
                    <div className="ProductKeyDetailColumn">
                        <h1>{productInfo.name}</h1>
                        <h2>{`${productInfo.price}`}</h2>
                        <Button startIcon={<EditIcon />} color="primary" variant="contained" onClick={() => {
                            ChangeCurrentOverlay("updateProduct")
                        }} size="large">Edit Details</Button>
                        <div style={{marginTop: "20px"}}>
                            <Button startIcon={<DeleteIcon />} color="error" variant="contained" onClick={() => {
                                ChangeCurrentOverlay("deleteConfirm")
                            }} size="large">Delete Product</Button>
                        </div>
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
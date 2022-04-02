import React, {useState, useEffect} from 'react';
import Header from '../components/Header';
import Footer from '../components/Footer';
import '../styles/sellerProductPage.css';

import gatorMug from '../images/gator.jpg';
import ArrowBackIcon from '@mui/icons-material/ArrowBack';
import Button from '@mui/material/Button';
import TextField from "@mui/material/TextField";
import Avatar from '@mui/material/Avatar';
import DeleteIcon from '@mui/icons-material/Delete';
import SentimentDissatisfiedIcon from '@mui/icons-material/SentimentDissatisfied';
import ShoppingCartIcon from '@mui/icons-material/ShoppingCart';

import settings from '../settings'
import '../styles/sellerStorePage.css';
import { useParams, Link } from 'react-router-dom';
import { useNavigate } from "react-router-dom";
import { useSearchParams } from "react-router-dom";

export default function ProductPage() {
    const navigate = useNavigate(); // to redirect using navigate()
    const [searchParams, setSearchParams] = useSearchParams(); // to get url parameter "?liveId=someLiveId" if it came from a livestream

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

                {currentOverlay === 'confirmPurchase' && (
                    <div class="go-live-overlay" >
                        <div class="transparentBG"/>
                            <div class="stream-link-container" style={{top: "35vh"}}>
                            
                            <div style={{textAlign: "center", paddingTop: "4px"}}>
                                <h2 style={{fontWeight: "normal", paddingBottom: "17px"}}>You have successfully purchased "{productInfo.name}"!</h2>
                                <Button variant="contained" color="primary" onClick={() => {
                                    ChangeCurrentOverlay("none");
                                }} size="large" style={{marginRight: "0"}}>Okay</Button>
                            </div>
                        </div>
                    </div>
                )}

                {currentOverlay === 'pleaseSignIn' && (
                    <div class="go-live-overlay" >
                        <div class="transparentBG"/>
                            <div class="stream-link-container" style={{top: "35vh"}}>
                            
                            <div style={{textAlign: "center", paddingTop: "4px"}}>
                                <h2 style={{fontWeight: "normal", paddingBottom: "17px"}}>Please sign-in to purchase a product.</h2>
                                <Button variant="contained" color="primary" onClick={() => {
                                    ChangeCurrentOverlay("none");
                                }} size="large" style={{marginRight: "0"}}>Okay</Button>
                            </div>
                        </div>
                    </div>
                )}
            </div>
        );
    }    

    function PurchaseProduct(quantity) {
        var jwtToken = window.sessionStorage.getItem("user-jwtToken");
        if (jwtToken === null || jwtToken === "") {
            ChangeCurrentOverlay("pleaseSignIn");
            return;
        }
        
        // Include LiveId on body of POST IF it is present in the url:
        if (searchParams.get("liveId") === null || searchParams.get("liveId") === "") { // liveId is not set
            const requestOptions = {
                method: 'POST',
                headers: {'Authorization': jwtToken},
                body: JSON.stringify({quantity: quantity})
            };
            fetch(settings.apiHostURL + `product/${productID}`, requestOptions)
            .then(response => response.json())
            .then(response => {
                if (response.status === 0) {
                    ChangeCurrentOverlay("confirmPurchase"); // display purchase confirmation
                    getProductInfo(); // re-get product info to see if it is now out of stock
                }
                else {
                    console.log("ERROR: Purchase product API did not respond with 'success' status code 0.");
                }
            })
            .catch((error) => {
                console.error(error);
            });
        } else { // liveId is set
            const requestOptions = {
                method: 'POST',
                headers: {'Authorization': jwtToken},
                body: JSON.stringify({quantity: quantity, liveId: searchParams.get("liveId")})
            };
            fetch(settings.apiHostURL + `product/${productID}`, requestOptions)
            .then(response => response.json())
            .then(response => {
                if (response.status === 0) {
                    ChangeCurrentOverlay("confirmPurchase"); // display purchase confirmation
                    getProductInfo(); // re-get product info to see if it is now out of stock
                }
                else {
                    console.log("ERROR: Purchase product API did not respond with 'success' status code 0.");
                }
            })
            .catch((error) => {
                console.error(error);
            });
        }
    }
    
    return (
        <div className="RootFlexContainer">
            <Header />

            <Overlay/>
            <div className="ProductDetailsRow">
                <div className="ProductImage">
                    <img src={gatorMug} alt="Gator Mug"/>
                </div>
                {productInfo ? 
                    <div className="ProductKeyDetailColumn">
                        <h1>{productInfo.name}</h1>
                        <h2>${productInfo.price}</h2>
                        <div style={{marginTop: "20px"}}>
                            {productInfo.quantity > 0 && (
                                <Button startIcon={<ShoppingCartIcon />} color="primary" variant="contained" onClick={() => {
                                    PurchaseProduct(1); //1 at a time for now...
                                }} size="large">Purchase Product</Button>
                            )}

                            {productInfo.quantity <= 0 && (
                                <Button startIcon={<SentimentDissatisfiedIcon />} disabled="true" color="primary" variant="contained" onClick={() => {
                                    PurchaseProduct(1); //1 at a time for now...
                                }} size="large">Out of Stock</Button>
                            )}
                        </div>
                    </div> 
                    : 
                    null
                }
                {productInfo ?
                    <div style={{flex: 1}} className="flexCenter">
                        <div className="ProductStoreInfo flexCenter colFlex">
                            <h1>{productInfo.storeId}</h1>
                            <Avatar sx={{ bgcolor: 'navy', fontSize: 50, width: 120, height: 120}}>UF</Avatar>
                            <Button component={Link} to={`/store/${productInfo.storeId}`} startIcon={<ArrowBackIcon />} color="secondary" variant="contained" size="large" sx={{marginBottom: 3}}>Back to Store</Button>
                            {/* <Button startIcon={<EditIcon />} color="primary" variant="contained" size="large" sx={{marginBottom: 1}}>Edit Store</Button> */}
                        </div>
                    </div>
                    :
                    null
                }
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
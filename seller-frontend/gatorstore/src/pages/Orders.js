import React, {useState, useEffect} from 'react';
import { Link } from 'react-router-dom';
import Header from '../components/Header';
import Footer from '../components/Footer';
import '../styles/orders.css';

import EditIcon from '@mui/icons-material/Edit';
import OpenInNewIcon from '@mui/icons-material/OpenInNew';
import Button from '@mui/material/Button';
import gatorPlush from '../images/gator-plush.png';

import settings from '../settings'
import { useNavigate } from "react-router-dom";

function MyOrders() {
    const navigate = useNavigate(); // to redirect using navigate()
    
    // UNCOMMENT WHEN USING SERVER:
    /* On load: Check that user is logged (if not, back to landing page they go)
    useEffect(() => {
        if (window.sessionStorage.getItem("user-jwtToken") === null) {
            alert("Please login to see your orders.");
            navigate("/");
        }

        // Get first page
        GetPage(0);
    }, []);*/

    function OrderEntry(orderInfo) {
        return (
            <div className="StoreEntry flexCenter rowFlex">
                <div style={{ flex: 0.7 }} className="flexCenter">
                    <img style={{height: "80px", width: "80px"}} src={gatorPlush} />
                </div>
                <div style={{ flex: 3 }}>
                    <h1 className="StoreEntryTitle">{orderInfo.productName} <OpenInNewIcon /></h1>
                    <div className="StoreEntryDetailRow">
                        <p><b>Quantity:</b> {orderInfo.quantity}</p>
                        <p>|</p>
                        <p><b>Store:</b> name (unless /orders/storeID)</p>
                        <p>|</p>
                        <p><b>ID:</b> {orderInfo.id}</p>
                    </div>
                </div>
                <div style={{ flex: 1}} className="flexCenter">
                <Button component={Link} to={'/product/' + orderInfo.id} startIcon={<EditIcon />} variant="contained" color="primary" sx={{ marginBottom: 1 }}>View Product</Button>
                    <p style={{fontSize: "20px", paddingTop: "8px"}}><b>Price:</b> ${orderInfo.subTotal}</p>
                </div>
            </div>
        );
    } 

    const [orderArray, SetOrderArray] = useState([{
        productName: "Test", productId: "25S", quantity: "3", subTotal: "35"
    }]);

    // For pagination:
    const [currProductPage, ChangeProductPage] = useState(0);
    var maxProductPage = 1; // default

    // Calls on GetPage() to get a new product page upon the user scrolling down.
    function ScrollDown() {
        // Only request more products if current page number is below max:
        if (currProductPage <= maxProductPage) {
            ChangeProductPage(currProductPage + 1);
            GetPage(currProductPage);
        }
    }

    function GetPage(pageNum) {
        /* Get JWT Token for POST request header:
        var jwtToken = window.sessionStorage.getItem("user-jwtToken");
        var userId = window.sessionStorage.getItem("user-id");
        
        // Call API to get product list:
        const requestOptions = {
        method: 'GET',
        headers: {
            'Authorization': jwtToken
        }
        };
        fetch(settings.apiHostURL + 'user/' + userId + '/order-list?page=' + pageNum, requestOptions)
        .then(response => response.json())
        .then(response => {
            if (response.status === 0) {
                // if page requested isn't more than max page: Add orders of this new page to "orderArray"
                if (pageNum <= response.result.maxPage && response.result.orderList != null) {
                    SetOrderArray(orderArray.concat(response.result.orderList));
                }
                // Set max page number so that this fetch isn't even called if it is an invalid page number
                maxProductPage = response.result.maxPage;
            }
            else {
            console.log("ERROR: Order Page API did not respond with 'success' status code.");
            }
        })
        .catch((error) => {
            console.error(error);
        });*/
    }

    function OrderList() {
        return (
            <div class="store-list-container" onScroll={(e) => {
                if (e.target.scrollHeight - e.target.scrollTop === e.target.clientHeight) {
                  ScrollDown();
                }
              }}>
                <div class="store-list-container">
                    <div className="storeListContainer">
                    {orderArray && orderArray.length > 0 && ( orderArray.map(function (order) {
                        return(
                            <OrderEntry productName={order.productName} id={order.productId} quantity={order.quantity} subTotal={order.subTotal}/>
                            );
                        })
                    )}
                    {orderArray.length == 0 && (
                        <div>- You have no orders. -</div>
                    )}
                    </div>
                </div>
            </div>
        );
    }

    return (
        <div className="RootFlexContainer">
        <Header />

        <div className="flexCenter colFlex">
            <div className="storeListSubHeader">
            <h1>Order History:</h1>
            </div>
        
            <OrderList />
        </div>
        <Footer />
        </div>
    );
}

export default MyOrders;
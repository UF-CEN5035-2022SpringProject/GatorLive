import React, {useState, useEffect} from 'react';
import { Link } from 'react-router-dom';
import Header from '../components/Header';
import Footer from '../components/Footer';
import '../styles/orders.css';

import OpenInNewIcon from '@mui/icons-material/OpenInNew';
import Button from '@mui/material/Button';
import settings from '../settings'
import { useNavigate } from "react-router-dom";
import { useParams } from 'react-router-dom';

import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableCell from '@mui/material/TableCell';
import TableContainer from '@mui/material/TableContainer';
import TableHead from '@mui/material/TableHead';
import TableRow from '@mui/material/TableRow';
import Paper from '@mui/material/Paper';
import SentimentDissatisfiedIcon from '@mui/icons-material/SentimentDissatisfied';

function Livestreams() {
    const navigate = useNavigate(); // to redirect using navigate()
    const { storeId } = useParams();
    
    // On load: Check that user is logged (if not, back to landing page they go)
    useEffect(() => {
        if (window.sessionStorage.getItem("user-jwtToken") === null) {
            alert("Please login to see your livestreams.");
            navigate("/");
        }

        // Get first page
        GetPage(0);
    }, []);

    function LivestreamEntry(stream) {
        return (
            <div className="LivestreamEntry rowFlex">
                <p className="LivestreamTitle"><b>{stream.date}</b> | "{stream.streamName}"</p>
                <p className="LivestreamTitle" style={{fontSize: "20px"}}><b>ID:</b> {stream.streamId}</p>
                <p className="LivestreamTitle" style={{fontSize: "20px"}}><b>Orders Made:</b></p>
                <TableContainer component={Paper} style={{maxHeight: 140,width: "98%", margin: "auto"}}>
                    <Table sx={{ minWidth: 650}} size="small" aria-label="a dense table" stickyHeader>
                        <TableHead>
                            <TableRow>
                                <TableCell>Product</TableCell>
                                <TableCell align="right">Quantity</TableCell>
                                <TableCell align="right">Subtotal ($)</TableCell>
                                <TableCell align="right">User</TableCell>
                            </TableRow>
                        </TableHead>

                        <TableBody>
                            {console.log(liveOrdersArray)}
                            {
                                liveOrdersArray && liveOrdersArray.length > 0 && liveOrdersArray[stream.index].map(function (order) {
                                    return(
                                    <TableRow key={order.name}>
                                        <TableCell component="th" scope="row">
                                            {order.name}
                                        </TableCell>
                                        <TableCell align="right">{order.quantity}</TableCell>
                                        <TableCell align="right">{order.subTotal}</TableCell>
                                        <TableCell align="right">{order.user}</TableCell>
                                    </TableRow>
                                    );
                                })
                                
                            }
                        </TableBody>
                    </Table>
                </TableContainer>
            </div>
        );
    }


    // Live-list!:
    const [liveList, SetLiveList] = useState([]);

    const testLiveOrders = [
        {name: "Gator Jacket", quantity: "2", subTotal: "35.50", user: "user1"},
        {name: "Gator Mug", quantity: "1", subTotal: "3.50", user: "user6"},
        {name: "Gator Laptop", quantity: "1", subTotal: "500.60", user: "user5"},
        {name: "Gator Table", quantity: "3", subTotal: "56.80", user: "user2"},
    ];

    // To see if store even exists:
    const [actualStore, SetIfActualStore] = useState(true);

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

    useEffect(() => {
        if (liveList != undefined) {
            liveList.forEach((livestream, index) => {
                liveOrdersArray.concat([]);
                GetLiveOrders(0, livestream.id, index);   
            })
        }
      }, [liveList]);

    function GetLiveOrders(pageNum, liveId, index) {
        // Get JWT Token for POST request header:
        var jwtToken = window.sessionStorage.getItem("user-jwtToken");
                        
        // Call API to get product list:
        const requestOptions = {
        method: 'GET',
        headers: {
            'Authorization': jwtToken
        }
        };
        fetch(settings.apiHostURL + 'live/' + liveId + '/order-list?page=' + pageNum, requestOptions)
        .then(response => response.json())
        .then(response => {
            if (response.status === 0) {
                if (response.result.orderList !== null) {
                    
                    //console.log(response.result.orderList)

                    //SetLiveOrdersArray(liveOrdersArray.concat(response.result.orderList));
                    liveOrdersArray[index] = response.result.orderList;
                    SetLiveOrdersArray(liveOrdersArray);
                } 
            }
            else {
                console.log("ERROR: Order Page API did not respond with 'success' status code.");
            }
        })
        .catch((error) => {
            console.error(error);
        });
    }

    // Array of arrays for live orders. Used to render by index
    const [liveOrdersArray, SetLiveOrdersArray] = useState([]);

    function GetPage(pageNum) {
        // Get JWT Token for POST request header:
        var jwtToken = window.sessionStorage.getItem("user-jwtToken");
        console.log("GetPage")
        // Call API to get product list:
        const requestOptions = {
        method: 'GET',
        headers: {
            'Authorization': jwtToken
        }
        };
        fetch(settings.apiHostURL + 'store/' + storeId + '/live-list?page=' + pageNum, requestOptions)
        .then(response => response.json())
        .then(response => {
            if (response.status === 0) {
                // if page requested isn't more than max page: Add orders of this new page to "orderArray"
                if (pageNum <= response.result.maxPage && response.result.liveList != null) {
                    SetLiveList(liveList.concat(response.result.liveList));
                    console.log("finished live-list")
                }
                // Set max page number so that this fetch isn't even called if it is an invalid page number
                maxProductPage = response.result.maxPage;
                SetIfActualStore(true)
            }
            else {
                console.log("ERROR: Order Page API did not respond with 'success' status code.");
                SetIfActualStore(false)
            }
        })
        .catch((error) => {
            console.error(error);
        });
    }

    function LivestreamList() {
        return (
            <div class="store-list-container">
                {actualStore && (
                    <div class="store-list-container" onScroll={(e) => {
                        if (e.target.scrollHeight - e.target.scrollTop === e.target.clientHeight) {
                        ScrollDown();
                        }
                    }}>
                        <div class="store-list-container">
                            <div className="storeListContainer">
                            {liveList && liveList.length > 0 && ( liveList.map(function (stream, index) {
                                return(
                                    <LivestreamEntry date={stream.createTime.substring(0, 10)} index={index} streamName={stream.title} streamId={stream.id}/>
                                    );
                                })
                            )}
                            {liveList.length == 0 && (
                                <div>- {storeId} has no past livestreams. -</div>
                            )}
                            </div>
                        </div>
                    </div>
                )}

                {!actualStore && (
                    <div className="storeListContainer">
                        <SentimentDissatisfiedIcon fontSize='large' color='warning'/>
                    </div>
                )}
            </div>
        );
    }

    return (
        <div className="RootFlexContainer">
            <Header />

            <div className="flexCenter colFlex">
                <div className="storeListSubHeader">
                    {actualStore && (
                        <h1>{storeId}'s Livestream History:</h1>
                    )}
                    {!actualStore && (
                        <p style={{fontSize: "30px"}}>You don't have access to <b>{storeId}</b></p>
                    )}
                </div>
            
                <LivestreamList />
            </div>
            <Footer />
        </div>
    );
}

export default Livestreams;
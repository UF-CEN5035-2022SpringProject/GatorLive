import React, {useState, useEffect} from 'react';
import { Link } from 'react-router-dom';
import Header from '../components/Header';
import Footer from '../components/Footer';
import '../styles/orders.css';

import OpenInNewIcon from '@mui/icons-material/OpenInNew';
import Button from '@mui/material/Button';
import settings from '../settings'
import { useNavigate } from "react-router-dom";

import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableCell from '@mui/material/TableCell';
import TableContainer from '@mui/material/TableContainer';
import TableHead from '@mui/material/TableHead';
import TableRow from '@mui/material/TableRow';
import Paper from '@mui/material/Paper';

function Livestreams() {
    const navigate = useNavigate(); // to redirect using navigate()
    
    // UNCOMMENT WHEN USING SERVER:
    /* On load: Check that user is logged (if not, back to landing page they go)
    useEffect(() => {
        if (window.sessionStorage.getItem("user-jwtToken") === null) {
            alert("Please login to see your livestreams.");
            navigate("/");
        }

        // Get first page
        GetPage(0);
    }, []);*/

    function LivestreamEntry(stream) {
        return (
            <div className="LivestreamEntry rowFlex">
                <p className="LivestreamTitle"><b>{stream.date}</b> | {stream.streamName}</p>
                <p className="LivestreamTitle" style={{fontSize: "22px"}}>Orders Made:</p>
                <TableContainer component={Paper} style={{width: "98%", margin: "auto"}}>
                    <Table sx={{ minWidth: 650 }} size="small" aria-label="a dense table">
                        <TableHead>
                            <TableRow>
                                <TableCell>Product</TableCell>
                                <TableCell align="right">Quantity</TableCell>
                                <TableCell align="right">Subtotal ($)</TableCell>
                                <TableCell align="right">User</TableCell>
                            </TableRow>
                        </TableHead>

                        <TableBody>
                            {testLiveOrders.map((order) => (
                                <TableRow key={order.name}>
                                    <TableCell component="th" scope="row">
                                        {order.name}
                                    </TableCell>
                                    <TableCell align="right">{order.quantity}</TableCell>
                                    <TableCell align="right">{order.subTotal}</TableCell>
                                    <TableCell align="right">{order.user}</TableCell>
                                </TableRow>
                            ))}
                        </TableBody>
                    </Table>
                </TableContainer>
            </div>
        );
    }


    // Live-list!:
    const [liveList, SetLiveList] = useState([{
        date: "4/9/22", streamName: "Hi guys", streamId: "49SKD2", store:"gatorstore-1"
    }]);

    const testLiveOrders = [
        {name: "Gator Jacket", quantity: "2", subTotal: "35.50", user: "user1"},
        {name: "Gator Mug", quantity: "1", subTotal: "3.50", user: "user6"},
        {name: "Gator Laptop", quantity: "1", subTotal: "500.60", user: "user5"},
        {name: "Gator Table", quantity: "3", subTotal: "56.80", user: "user2"},
    ];

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
                    SetLiveList(liveList.concat(response.result.orderList));
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

    function LivestreamList() {
        return (
            <div class="store-list-container" onScroll={(e) => {
                if (e.target.scrollHeight - e.target.scrollTop === e.target.clientHeight) {
                  ScrollDown();
                }
              }}>
                <div class="store-list-container">
                    <div className="storeListContainer">
                    {liveList && liveList.length > 0 && ( liveList.map(function (stream) {
                        return(
                            <LivestreamEntry date={stream.date} streamName={stream.streamName} streamId={stream.streamId}/>
                            );
                        })
                    )}
                    {liveList.length == 0 && (
                        <div>- You have no past livestreams. -</div>
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
            <h1>Past Livestream History:</h1>
            </div>
        
            <LivestreamList />
        </div>
        <Footer />
        </div>
    );
}

export default Livestreams;
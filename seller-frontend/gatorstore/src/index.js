import React from 'react';
import ReactDOM from 'react-dom';
import {
  BrowserRouter,
  Routes,
  Route
} from "react-router-dom";

import SellerStoreList from './pages/SellerStoreList'
import ProductPage from './pages/ProductPage'
import LandingPage from './pages/LandingPage'
import SellerStorePage from './pages/SellerStorePage';
import LoginRedirect from './pages/LoginRedirect'
import Orders from './pages/Orders'
import PastLivestreams from './pages/Livestreams'

ReactDOM.render(
  <BrowserRouter basename="/">
    <Routes>
      <Route path="/store-list" element={<SellerStoreList/>} />
      <Route path="/product/:productID" element={<ProductPage/>} />
      <Route path="/" element={<LandingPage/>} />
      <Route path="/store/:storeID" element={<SellerStorePage/>} />
      <Route path="/login" element={<LoginRedirect/>} />

      <Route path="/:storeId/orders" element={<Orders/>} /> 
      <Route path="/:storeId/livestreams" element={<PastLivestreams/>} /> 
    </Routes>
  </BrowserRouter>,
  document.getElementById('root')
);


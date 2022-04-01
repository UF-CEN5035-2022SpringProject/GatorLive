import React from 'react';
import ReactDOM from 'react-dom';
import {
  BrowserRouter,
  Routes,
  Route
} from "react-router-dom";

import ProductPage from './pages/ProductPage'
import StorePage from './pages/Store'
import LandingPage from './pages/LandingPage'
import LoginRedirect from './pages/LoginRedirect'

ReactDOM.render(
  <BrowserRouter basename="/">
    <Routes>
      <Route path="/" element={<LandingPage/>} />
      <Route path="/product/:productID" element={<ProductPage/>} />
      <Route path="/store/:storeID" element={<StorePage/>} />
      <Route path="/login" element={<LoginRedirect/>} />
    </Routes>
  </BrowserRouter>,
  document.getElementById('root')
);


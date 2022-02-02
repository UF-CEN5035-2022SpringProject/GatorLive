import React from 'react';
import ReactDOM from 'react-dom';
import {
  BrowserRouter,
  Routes,
  Route
} from "react-router-dom";

import SellerStoreList from './pages/SellerStoreList'
import ProductPage from './pages/ProductPage'

ReactDOM.render(
  <BrowserRouter>
    <Routes>
      <Route path="/store-list" element={<SellerStoreList/>} />
      <Route path="/product-page" element={<ProductPage/>} />
    </Routes>
  </BrowserRouter>,
  document.getElementById('root')
);


import React from 'react';
import ReactDOM from 'react-dom';
<<<<<<< Updated upstream
import './index.css';
import App from './App';
import reportWebVitals from './reportWebVitals';

ReactDOM.render(
  <React.StrictMode>
    <App />
  </React.StrictMode>,
  document.getElementById('root')
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
=======
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

>>>>>>> Stashed changes

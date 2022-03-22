import React from 'react';
import ReactDOM from 'react-dom';
import {
  BrowserRouter,
  Routes,
  Route
} from "react-router-dom";

import StorePage from './pages/Store'
import LandingPage from './pages/LandingPage'

ReactDOM.render(
  <BrowserRouter basename="/">
    <Routes>
      <Route path="/" element={<LandingPage/>} />
      <Route path="/store" element={<StorePage/>} />
    </Routes>
  </BrowserRouter>,
  document.getElementById('root')
);


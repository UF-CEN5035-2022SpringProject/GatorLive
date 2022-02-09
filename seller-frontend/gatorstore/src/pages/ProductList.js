import React from 'react';
import Header from '../components/Header';
import Footer from '../components/Footer';
import {Grid} from "@material-ui/core";
import Productcontent from '../components/ProductContent';
function ProductList() {
  return (
  <div>
  <div>
  <Header/>
  </div> 
  <Grid container direction='column'>
       <Grid item container>
         <Grid item xs={false} sm={2} />
         <Grid item xs={12} sm={8}>
            <Productcontent/>
         </Grid>
         <Grid item xs={false} sm={2} />
       </Grid>
    </Grid>     
  <div>
    <Footer/>
  </div>

 </div>
  );
}

export default ProductList;

  
import React from 'react';
import Header from '../components/Header.js';
import Footer from '../components/Footer';
import ActionAreaCard from'../components/aboutcard';
import {Grid} from "@material-ui/core";
import ImageSlider from '../components/ImageSlider';
import { SliderData } from '../components/SliderData';




function landingpage() {
    
  return( 
  <div>
      <div >
          <Header/>
      </div>
      <div>
          <ImageSlider slides={SliderData} />
      </div>
      <div>
          <h3>About Team</h3>
      <Grid container>
          <Grid item xs={3}>
          <ActionAreaCard/>
          </Grid>
          <Grid item xs={3}>
          <ActionAreaCard/>
          </Grid>
          <Grid item xs={3}>
          <ActionAreaCard/>
          </Grid>
          <Grid item xs={3}>
          <ActionAreaCard/>
          </Grid>
      </Grid>
      </div>
      <div>
         <Footer/>
     </div>

  </div>
  );
}

export default landingpage;


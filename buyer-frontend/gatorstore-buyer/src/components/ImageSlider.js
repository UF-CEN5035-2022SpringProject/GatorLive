import React, {useState} from 'react';
import { FaArrowAltCircleRight, FaArrowAltCircleLeft } from 'react-icons/fa';

import gator from '../images/slider1.png'
import gator2 from '../images/slider2.png'
import gator3 from '../images/slider3.png'

const ImageSlider = ({slides}) => {
    const [current, setCurrent] = useState(0)
    const length = 3
  

    const nextSlide = () => {
        setCurrent(current === length - 1 ? 0 : current + 1);
      };
    
      const prevSlide = () => {
        setCurrent(current === 0 ? length - 1 : current - 1);
      };

  

    if (!Array.isArray(slides) || slides.length <= 0) {
        return null;
      }
   
    return( 
      <section className='slider'> 
        <FaArrowAltCircleLeft className='left-arrow' onClick={prevSlide}/>
        <FaArrowAltCircleRight className='right-arrow' onClick={nextSlide}/>
     {/* {SliderData.map((slide, index) => {

         return( */}
        <div className={current === 0 ? 'slide active' : 'slide'} key={0}>
          {current === 0 &&( <img src={gator} alt="Promoted Image" className='image'/> )}  
        </div>

        <div className={current === 1 ? 'slide active' : 'slide'} key={1}>
          {current === 1 &&( <img src={gator2} alt="Promoted Image" className='image'/> )}  
        </div>

        <div className={current === 2 ? 'slide active' : 'slide'} key={2}>
          {current === 2 &&( <img src={gator3} alt="Promoted Image" className='image'/> )}  
        </div>
         {/* )
     } )} */}
     </section>
    
    )
};

export default ImageSlider;

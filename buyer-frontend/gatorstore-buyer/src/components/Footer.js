import React from 'react';
import LinkedInIcon from '@mui/icons-material/LinkedIn';

export default function Footer() {
  return (
    <div className="footer flexCenter">
      <div style={{width: '80%', height: '100%', display: 'flex', flexDirection: 'column'}}>
        <div className='footer-title' style={{width: '100%', height: 100}}> 
          About Us 
        </div>
        <div style={{width: '100%', height: 250, display: 'flex', flexDirection: 'row', justifyContent: 'space-between'}}>
          <div className='footer-block'>
            <div className='footer-name'>
              Yi-Ming Chang
              <a 
                href='https://www.linkedin.com/in/yiming-chang/'
                className='footer-linkedIn-btn'
                target='_blank'
              >
                <LinkedInIcon/>
              </a>
            </div>
            <div className='footer-member-content'>
              Project Manager <br/>
              Full-Stack Developer <br/>
              [ Golang, React ] <br/>
            </div>
          </div>
          <div className='footer-block'>
            <div className='footer-name'>
              Hung-You Chou
              <a 
                href='https://www.linkedin.com/in/hung-you-chou-039811160/'
                className='footer-linkedIn-btn'
                target='_blank'
              >
                <LinkedInIcon/>
              </a>
            </div>
            <div className='footer-member-content'>
              Backend Developer <br/>
              [Golang, NoSQL] <br/>
            </div>
          </div>
          <div className='footer-block'>
            <div className='footer-name'>
              Sebastian Llerena
              <a 
                href='https://www.linkedin.com/in/sebastian-llerena/'
                className='footer-linkedIn-btn'
                target='_blank'
              >
                <LinkedInIcon/>
              </a>
            </div>
            <div className='footer-member-content'>
              Frontend Developer <br/>
              [ React, Hook, Css ] <br/>
            </div>
          </div>
        </div>
        <div style={{textAlign: 'left'}}>GatorStore © 2022</div>
      </div>
    </div>
  );
}

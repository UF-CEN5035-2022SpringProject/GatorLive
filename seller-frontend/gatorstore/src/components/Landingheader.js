import * as React from 'react';
import ReactDOM from 'react-dom';
import '../styles/app.css';
import TextField from "@mui/material/TextField";
import Button from '@mui/material/Button';
import SearchIcon from '@mui/icons-material/Search';
import PersonIcon from '@mui/icons-material/Person';
import Menu from '@mui/material/Menu';
import MenuItem from '@mui/material/MenuItem';
import {Link} from 'react-router-dom'
import { makeStyles } from '@material-ui/core/styles';
import Login from'./Login'

const useStyles = makeStyles({
  searchBar: {
    backgroundColor: 'white',
    borderRadius: '5px',
    width: 300
  }
})

function SearchBar() {
  const classes = useStyles();
  return <TextField className={classes.searchBar} variant ="outlined" color="primary" placeholder="Search" size="small"/>
}

export default function Header() {
  const [anchorEl, setAnchorEl] = React.useState(null); // hook for the user menu dropdown
  const open = Boolean(anchorEl);

  const handleOpenMenu = e => {
    setAnchorEl(e.currentTarget); // tells you which element has been clicked in the menu
  };
  const handleClose = () => {
    setAnchorEl(null);
  };

  return (<div className="header">
      <div className="headerLogo flexCenter">GatorStore</div>
      <div className="searchBarContainer flexCenter">
        <SearchBar />
        <Button color="primary" variant="contained" size="medium"><SearchIcon/></Button>
      </div>
      <div className="googlebtn">
     
  <React.StrictMode>
    <div className="g-signin">
     <a href='https://accounts.google.com/o/oauth2/auth?access_type=offline&client_id=138444517704-gg6649ok973letdlh55bpte8bna7721o.apps.googleusercontent.com&redirect_uri=http%3A%2F%2Flocalhost:3000%2Ftesting.html&response_type=code&scope=https%3A%2F%2Fwww.googleapis.com%2Fauth%2Fyoutube+https%3A%2F%2Fwww.googleapis.com%2Fauth%2Fuserinfo.email+https%3A%2F%2Fwww.googleapis.com%2Fauth%2Fuserinfo.profile&state=state`' >
       <button className='login'>
         Sign In 
       </button>
       </a>
    </div>
  </React.StrictMode>
  

      </div>
  </div>);
}

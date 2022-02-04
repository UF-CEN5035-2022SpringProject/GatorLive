import * as React from 'react';
import { useNavigate } from 'react-router-dom';

import '../styles/app.css';
import TextField from "@mui/material/TextField";
import Button from '@mui/material/Button';
import HomeIcon from '@mui/icons-material/Home';
import StorefrontIcon from '@mui/icons-material/Storefront';
import FormatListBulletedIcon from '@mui/icons-material/FormatListBulleted';
import SearchIcon from '@mui/icons-material/Search';
import PersonIcon from '@mui/icons-material/Person';
import PushPinIcon from '@mui/icons-material/PushPin';
import LogoutIcon from '@mui/icons-material/Logout';
import Menu from '@mui/material/Menu';
import MenuItem from '@mui/material/MenuItem';

import { makeStyles } from '@material-ui/core/styles';
import { Navigate } from 'react-router-dom';

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

function SignedInDropdown(userData) {
  const navigate = useNavigate();

  const [anchorEl, setAnchorEl] = React.useState(null); // hook for the user menu dropdown
  const open = Boolean(anchorEl);

  const handleOpenMenu = e => {
    setAnchorEl(e.currentTarget); // tells you which element has been clicked in the menu
  };
  const handleClose = () => {
    setAnchorEl(null);
  };
  const SignOut = () => {
    window.sessionStorage.clear();
    window.location.href = "http://localhost:3000/landingpage";
  };

  return (<>
    <Button         
    color="primary" variant="contained" size="medium"
    startIcon={<PersonIcon />}
    id="basic-button"
    aria-controls={open ? 'basic-menu' : undefined}
    aria-haspopup="true"
    aria-expanded={open ? 'true' : undefined}
    onClick={handleOpenMenu}
    >
    {userData.name}
    </Button>
    <Menu
      id="basic-menu"
      anchorEl={anchorEl}
      open={open}
      onClose={handleClose}
    >
      <div style={{padding: 10, color: 'blue'}}>{userData.email}</div>
      <MenuItem component="a" href="/landingpage"><HomeIcon style={{marginRight: 20}}/> Home</MenuItem>
      <MenuItem component="a" href="/store-list"><StorefrontIcon style={{marginRight: 20}}/> My Stores</MenuItem>
      <MenuItem component="a" href="/productlist"><FormatListBulletedIcon style={{marginRight: 20}}/> My Listings</MenuItem>
      <MenuItem component="a" href="/product-page"><PushPinIcon style={{marginRight: 20}}/> Pinned Listing</MenuItem>
      <MenuItem onClick={SignOut}><LogoutIcon style={{marginRight: 20}}/>Logout</MenuItem>
    </Menu>
  </>);
}

function GoogleButton() { // for when user is NOT signed in
  return (<div>- Put Google Button Here -</div>);
}

function LoadSessionActions() {
  // User has NOT signed in:
  if (sessionStorage.getItem('user-name') === null) {
    return (<GoogleButton />);
  }
  else { // User is signed in:
    let name = window.sessionStorage.getItem('user-name');
    let email = window.sessionStorage.getItem('user-email');
    let id = window.sessionStorage.getItem('user-id');
    let jwtToken = window.sessionStorage.getItem('user-jwtToken');

    return (<SignedInDropdown name={name} email={email}/>);
  }
}

export default function Header() {
  return (<div className="header">
      <div className="headerLogo flexCenter">GatorStore</div>
      <div className="searchBarContainer flexCenter">
        <SearchBar />
        <Button color="primary" variant="contained" size="medium"><SearchIcon/></Button>
      </div>
      <div className="accountButton flexCenter">
        <LoadSessionActions />
      </div>
  </div>);
}

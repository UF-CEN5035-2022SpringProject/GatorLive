import * as React from 'react';

import '../styles/app.css';
import TextField from "@mui/material/TextField";
import Button from '@mui/material/Button';
import SearchIcon from '@mui/icons-material/Search';
import PersonIcon from '@mui/icons-material/Person';
import Menu from '@mui/material/Menu';
import MenuItem from '@mui/material/MenuItem';

import { makeStyles } from '@material-ui/core/styles';

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
      <div className="accountButton flexCenter">
        <Button         
          color="primary" variant="contained" size="medium"
          startIcon={<PersonIcon />}
          id="basic-button"
          aria-controls={open ? 'basic-menu' : undefined}
          aria-haspopup="true"
          aria-expanded={open ? 'true' : undefined}
          onClick={handleOpenMenu}
        >
          My Account
        </Button>
        <Menu
          id="basic-menu"
          anchorEl={anchorEl}
          open={open}
          onClose={handleClose}
        >
          <MenuItem onClick={handleClose}>My Stores</MenuItem>
          <MenuItem onClick={handleClose}>Active Listings</MenuItem>
          <MenuItem onClick={handleClose}>Settings</MenuItem>
          <MenuItem onClick={handleClose}>Logout</MenuItem>
        </Menu>
      </div>
  </div>);
}

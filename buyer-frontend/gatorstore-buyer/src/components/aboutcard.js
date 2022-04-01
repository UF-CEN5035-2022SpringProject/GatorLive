import * as React from 'react';
import Card from '@mui/material/Card';
import CardContent from '@mui/material/CardContent';
import CardMedia from '@mui/material/CardMedia';
import Typography from '@mui/material/Typography';
import { CardActionArea } from '@mui/material';

const ActionAreaCard = props => {
  const { storeName } = props;
  return (
    <Card sx={{ maxWidth: 345 }}>
    <CardActionArea>
      <CardMedia
        component="img"
        height="300"
        image="https://m.media-amazon.com/images/I/61SNVLGQCCL._AC_SX679_.jpg"
        alt="store"
      />
      <CardContent>
        <Typography gutterBottom variant="h5" component="div">
          {storeName}
        </Typography>
        <Typography variant="body2" color="text.secondary">
          Body containing Details about the store.
        </Typography>
      </CardContent>
    </CardActionArea>
    </Card>
  );
};

export default ActionAreaCard;
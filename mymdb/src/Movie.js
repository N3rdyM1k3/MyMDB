import { Card, Grid, CardMedia, Typography } from '@mui/material';
import CardContent from "@mui/material/Card"
import { useTheme } from '@mui/material/styles'


function Movie(movie) {
  const theme = useTheme();

  return (
    <Grid item xs={12} sm={6} lg={4}>
      <Card sx={{ display: "flex", border: 4, borderRadius: 2, borderColor: 'secondary.dark' }}>
        <CardMedia component="img" image={movie.Poster} sx={{ width: 100, height: 150 }} />
        <CardContent sx={{ flex: '1 0 auto' }}>
          <Typography component="div" variant="h5">{movie.Title}</Typography>
          <Typography variant="p" component="div">{movie.Country}</Typography>
        </CardContent>
      </Card>
    </Grid>
  );
}

export default Movie;
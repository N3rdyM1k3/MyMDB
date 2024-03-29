import { Card, CardActions, Button, Grid, CardMedia, Typography, Box, Stack, Paper } from '@mui/material';
import CardContent from "@mui/material/Card"
import { useTheme } from '@mui/material/styles'
import { Link } from 'react-router-dom';


function Movie(movie) {
  const theme = useTheme();

  return (
    <Grid item xs={12} sm={6} lg={4}>
      <Card sx={{ display: "flex", border: 4, borderRadius: 2, borderColor: 'secondary.light' }}>
        
          <CardMedia component="img" image={movie.Poster} sx={{ width: 100, height: 150 }} />
            <Stack>
              <CardContent sx={{ p: 1, boxShadow:0 }}>
                    <MovieDetail {...movie} />
              </CardContent>
              <CardActions>
                  <Box sx={{ display:'flex', justifyContent: 'flex-end' }}>
                    <Button variant="outlined" color="primary" size="small" href={"../movies/"+movie.imdbID}>
                      <Link to={"../movies/"+movie.imdbID} state={{movie: movie}}>MORE</Link>
                    </Button>
                  </Box>
                  <Box sx={{ display:'flex', justifyContent: 'flex-end' }}>
                    <Button variant="outlined" color="primary" size="small">ACTION</Button>
                  </Box>
              </CardActions>
            </Stack>
      </Card>
    </Grid>
  );
}

function MovieDetail(movie) {
  return (
    <Box >
      <Typography component="div" variant="h5">{movie.Title}</Typography>
      <Typography variant="p" component="div">{movie.Country}  {movie.Year}</Typography>
      <Typography variant="p" component="div">{movie.Runtime} {movie.Rated}</Typography>
      <Typography variant="p" component="div">{movie.Genre}</Typography>
    </Box>
  );
}

export default Movie;
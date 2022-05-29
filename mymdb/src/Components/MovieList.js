import { Grid, Box } from '@mui/material';
import Movie from './Movie';

function MovieList(props) {
  const movies = props.movies;


  return (
    <>
    <Box sx={{ height: 50 }} />
    <Grid container>
      <Grid item xs={1}/>
      <Grid item xs={10}>
        <Grid container spacing={7}>
          {
            movies.map((m) => { return (<Movie {...m} />) })
          }
        </Grid>
      </Grid>
      <Grid item xs={1}/>
    </Grid></>
  );
}

export default MovieList;
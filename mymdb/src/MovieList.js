import {Grid} from '@mui/material';
import Movie from './Movie';

function MovieList(props) {

  const {movies} = {...props}; 
  return (
      <Grid container spacing={2}>
        {
          movies.map((m) => {return (<Movie {...m} />)})
        }
      </Grid>
  );
}

export default MovieList;
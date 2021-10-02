import {Grid} from '@mui/material';
import Movie from './Movie';

function MovieList(props) {

  const {movies, theme} = {...props}; 
  return (
      <Grid theme={theme} container spacing={2}>
        {
          movies.map((m) => {return (<Movie {...m} />)})
        }
      </Grid>
  );
}

export default MovieList;
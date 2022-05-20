import {Paper} from '@mui/material'
import { useParams, useNavigate, useLocation } from 'react-router-dom';
import {getMovie} from '../scripts/movies.js';

function MovieDetail(){

    let location = useLocation();
    let params = useParams();
    // TODO: Consider moving this logic out of presentation
    let movie;
    if (location.state && location.state.movie){
        movie = location.state.movie;
    }
    else{
        movie = getMovie(params.imdbID);
    }
    return (
        <Paper>
            <p>{movie.Title}</p>
        </Paper>
    );
}

export default MovieDetail;

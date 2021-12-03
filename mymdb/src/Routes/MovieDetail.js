import {Paper} from '@mui/material'
import { useParams, useNavigate } from 'react-router-dom';
import {getMovie} from '../data/movies.js';

function MovieDetail(){
    let params = useParams();
    let movie = getMovie(params.imdbID);
    return (
        <Paper>
            <p>{movie.Title}</p>
        </Paper>
    );
}

export default MovieDetail;

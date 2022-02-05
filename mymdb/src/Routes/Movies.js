import {getMovies} from '../scripts/movies.js';
import MovieList from '../Components/MovieList';

function Movies(){
    const movies = getMovies();
    return (<><MovieList movies={movies}/></>);
}

export default Movies;
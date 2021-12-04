import movies from './movies.json';


export function getMovie(imdbID) {
    return movies.find((movie) => movie.imdbID === imdbID);
  }

  export function getMovies() {
    return movies;
  }

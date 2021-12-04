import movies from './movies.json';


export function getMovie(imdbID) {
    return movies.find((movie) => movie.imdbID === imdbID);
  }

  export function getMovies() {
    return movies;
  }

export function searchAPI(title){
  return fetch("https://www.omdbapi.com/?s="+title+"&page=1&type=movie&apikey="+process.env.REACT_APP_API_KEY)
  .then(res => res.json())
  .then(res => {return res.Search})
}

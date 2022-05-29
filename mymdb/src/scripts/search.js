
export function searchAPI(title){
    return fetch("https://www.omdbapi.com/?s="+title+"&page=1&type=movie&apikey="+process.env.REACT_APP_API_KEY)
    .then(res => res.json())
    .then(res => {return res.Search})
  }

export function getMovie(imdbID) {
    return fetch("https://www.omdbapi.com/?i="+imdbID+"&page=1&type=movie&apikey="+process.env.REACT_APP_API_KEY)
    .then(res => res.json())
    .then(res => {return res.Search})
  }


  
import { useEffect, useState } from 'react'
import { getMovies } from '../data/movies';
import MovieList from '../Components/MovieList'

function Search(){
    const [results, setResults] = useState([]);

    useEffect(() => {
        getResults().then(setResults);
    }, [])

    return (<><MovieList movies={results} /></>);
}

export default Search;


async function getResults() {
        debugger;
        // Mimic fetching results 
        return fetch("https://www.omdbapi.com/?s=matrix&page=1&type=movie&apikey="+process.env.REACT_APP_API_KEY)
            .then(res => res.json())
            .then(res => {return res.Search})
            
        // TODO: Replace with search function  
}

function sleep(ms) {
    return new Promise(resolve => setTimeout(resolve, ms));
  }
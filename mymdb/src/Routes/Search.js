import { useEffect, useState } from 'react'
import { getMovies } from '../data/movies';
import MovieList from '../Components/MovieList'

function Search(){
    const [results, setResults] = useState([]);

    useEffect(() => {
        getResults().then(setResults)
    }, [])

    return (<><MovieList movies={results} /></>);
}

export default Search;


async function getResults() {
        // Mimic fetching results 
        await sleep(2000); 
        return getMovies();
        // TODO: Replace with search function  
}

function sleep(ms) {
    return new Promise(resolve => setTimeout(resolve, ms));
  }
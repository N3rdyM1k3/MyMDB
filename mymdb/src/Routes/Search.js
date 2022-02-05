import { useEffect, useState } from 'react'
import { searchAPI } from '../scripts/search';
import MovieList from '../Components/MovieList'

function Search(){
    const [results, setResults] = useState([]);

    useEffect(() => {
        searchAPI("Now and Then").then(setResults);
    }, [])

    return (<><MovieList movies={results} /></>);
}

export default Search;
import { useEffect, useState } from 'react'
import { searchAPI } from '../scripts/search';
import MovieList from '../Components/MovieList'
import { TextField } from '@mui/material';
import { useTheme } from '@mui/material/styles'

function Search(){
    const [results, setResults] = useState([]);
    const [searchTitle, setSearchTitle] = useState("");
    const theme = useTheme();

    useEffect(() => {
        searchAPI("Now and Then").then(setResults);
    }, [])

    const handleSubmit = evt => {
        evt.preventDefault();
        console.log(searchTitle)
        searchAPI(searchTitle).then(setResults);
    }

    const handleSearchChange = evt => {
        setSearchTitle(evt.target.value);
    }

    return (<>
    <form onSubmit={handleSubmit}>
        <TextField id="movie-title" label="Title" variant="filled" fullWidth sx={{ backgroundColor: 'background.paper' }} onChange={handleSearchChange} />
    </form>
    <MovieList movies={results} />
    </>);
}

export default Search;
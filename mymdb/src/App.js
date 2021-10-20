//import './App.css';
import movies from './data/movies.json';
import MovieList from './MovieList';
import {themeOptions } from './ThemeOptions';
import { ThemeProvider, createTheme } from '@mui/material/styles'
import { Paper } from '@mui/material';


const theme = createTheme({...themeOptions});

function App() {
  

  return (
    <div className="App">
      <ThemeProvider theme={theme} >
        <Paper>
          <MovieList movies={movies} />
        </Paper>
      </ThemeProvider>
    </div>
  );
}

export default App;

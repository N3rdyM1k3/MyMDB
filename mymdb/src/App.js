//import './App.css';
import movies from './data/movies.json';
import MovieList from './MovieList';
import { useTheme } from '@mui/material/styles'
import {themeOptions } from './ThemeOptions';
import { ThemeProvider } from '@mui/material/styles'


function App() {
  return (
    <div className="App">
      <ThemeProvider theme={themeOptions} >
        <MovieList movies={movies} />
      </ThemeProvider>
    </div>
  );
}

export default App;

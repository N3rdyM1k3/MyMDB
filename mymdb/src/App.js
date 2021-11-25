//import './App.css';
import movies from './data/movies.json';
import MovieList from './MovieList';
import {themeOptions } from './ThemeOptions';
import { ThemeProvider, createTheme } from '@mui/material/styles'
import CssBaseline from "@mui/material/CssBaseline";
import { Paper } from '@mui/material';


const theme = createTheme({...themeOptions});

function App() {
  

  return (

      <ThemeProvider theme={theme} >
          <CssBaseline />
          <MovieList movies={movies} />
      </ThemeProvider>
  );
}

export default App;

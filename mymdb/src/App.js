//import './App.css';
import movies from './data/movies.json';
import MovieList from './MovieList';
import { useTheme } from '@mui/material/styles'


function App() {
  const theme = useTheme();
  return (
    <div className="App">
      <MovieList theme={theme} movies={movies} />
    </div>
  );
}

export default App;

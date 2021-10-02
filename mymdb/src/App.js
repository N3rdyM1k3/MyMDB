import logo from './logo.svg';
import './App.css';
import {Grid, Card} from '@mui/material';
import movies from './data/movies.json';


function App() {
  return (
    <div className="App">
      <Grid>{movies.map((m) => {return (<Card>{m.Title}</Card>)})}</Grid>
    </div>
  );
}

export default App;

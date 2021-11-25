//import './App.css';
import MovieList from './MovieList';
import {themeOptions } from './ThemeOptions';
import { ThemeProvider, createTheme } from '@mui/material/styles'
import CssBaseline from "@mui/material/CssBaseline";
import { Paper } from '@mui/material';
import { Outlet } from 'react-router';


const theme = createTheme({...themeOptions});

function App() {
  

  return (

      <ThemeProvider theme={theme} >
          <CssBaseline />
          <Outlet />
      </ThemeProvider>
  );
}

export default App;

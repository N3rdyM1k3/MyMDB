//import './App.css';
import MovieList from './Components/MovieList';
import {themeOptions } from './ThemeOptions';
import { ThemeProvider, createTheme } from '@mui/material/styles'
import CssBaseline from "@mui/material/CssBaseline";
import { Box, Container } from '@mui/material';
import { Outlet } from 'react-router';


const theme = createTheme({...themeOptions});

function App() {
  

  return (

      <ThemeProvider theme={theme} >
          <CssBaseline />
          <Box sx={{ height: 100 }} /> {/* TODO: This is a place holder for where the menu will live */}
          <Container >
          <Outlet />
          </Container>
      </ThemeProvider>
  );
}

export default App;

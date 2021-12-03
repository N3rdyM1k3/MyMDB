import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import App from './App';
import reportWebVitals from './reportWebVitals';
import { BrowserRouter, Routes, Route } from 'react-router-dom';
// import MoviesList from './Components/MovieList'; // TODO: Change this to ./Routes/... as soon as you split MovieList
import MovieDetail from './Routes/MovieDetail';
import Movies from './Routes/Movies';


ReactDOM.render(
  <BrowserRouter>
  <Routes>
    <Route path="/" element={<App />}>
      <Route path="/movies" element={<Movies />} />
      <Route path="/movies/:imdbID" element={<MovieDetail />} />
    </Route>
    </Routes>
  </BrowserRouter>,
  document.getElementById('root')
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();

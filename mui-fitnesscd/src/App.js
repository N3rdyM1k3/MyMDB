import React, { useState } from 'react';
import Typography from '@material-ui/core/Typography'

function App() { 
  
  const [state, setState] = useState({ exercises: [], title: '' });
    // return (<Typography variant='display1' align='center' gutterBottom>        Exercises      </Typography>);
    return(<h1>Exercises</h1>)

}

export default App


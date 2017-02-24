import axios from 'axios';
import React, { Component } from 'react';
import logo from './logo.svg';
import './App.css';

class App extends Component {
  render() {
    return (
      <div className="App">
        <div className="App-header">
          <img src={logo} className="App-logo" alt="logo" />
          <h2>Welcome to React</h2>
        </div>
        <p className="App-intro">
          To get started, edit <code>src/App.js</code> and save to reload.
        </p>
        <button onClick={()=>{
          console.log("Hi John")
          axios.post('http://10.254.25.90:8000/register/').then((r) => {console.log(r.data)})}
        }>
        Hello</button>
      </div>
    );
  }
}

export default App;

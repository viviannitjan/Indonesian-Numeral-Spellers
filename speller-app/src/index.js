import React from 'react';
import ReactDOM from 'react-dom';
import App from './App';

ReactDOM.render(<App />, document.getElementById('root'));

document.getElementById("read").addEventListener('click', function(event){
    event.preventDefault();
});

document.getElementById('spell').addEventListener('click', function(e){
    e.preventDefault();
});
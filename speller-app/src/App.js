import React from 'react';
import ReactDOM from 'react-dom';
import axios from 'axios';
import './App.css';

class App extends React.Component{
    constructor(props){
        super(props);

    }

    handleSpell(){
        axios.get('http://localhost:8081/spell?number='+document.getElementById("input").value)
        .then(res => document.getElementById("outputbox").innerHTML = res.data.text);
    }

    handleRead(){
        var resultElement = document.getElementById("input").value;
        axios.post('http://localhost:8081/read', {text : resultElement})
        .then(res => document.getElementById("outputbox").innerHTML = res.data.number);
    }
    
    render(){
        return (
            <div class = "main">
                <div className="textbox">
                    <input id = "input" type = "text" />
                    <button className = "spell" id="spell" onClick={this.handleSpell}>Spell</button>
                    <button className = "read" id="read" onClick={this.handleRead}>Read</button>
                </div>

                <div className="outputbox" id="outputbox"></div>
            </div>
        )
    }
} 
export default App;

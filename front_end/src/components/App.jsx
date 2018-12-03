import React, {Component} from 'react'
import '../styles/App.css'

class App extends Component {

    render() {
        fetch ("http://localhost:8002/help").then(response => console.log(response));
        return (
          <div className="App-title">
              <h3> Welcome To RAudio. </h3>
          </div>
        )
    }
}

export default App;
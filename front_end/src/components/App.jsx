import React, {Component} from 'react'
import '../styles/App.css'
import createMuiTheme from "@material-ui/core/es/styles/createMuiTheme";
import orange from "@material-ui/core/es/colors/orange";
import MusicList from "./MusicList";

const lightTheme = createMuiTheme({
    palette: {
        primary: orange,
        secondary: {main: "#f3f4f5"}
    }
});

class App extends Component {

    constructor(props) {
        super(props);
        this.state = {
            list: [],
            alreadyCalled: false
        };
    }

    addItem(item) {
        let listx = this.state.list;
        listx.push(item);
        this.setState({list: listx})
    }

    getList() {
        fetch("http://localhost:8002/search?q=Ruth%20B").then(response => {
            response.json().then(json => {
                console.log(json);
                for (let i = 0; i < json.length; i++) {
                    this.addItem(json[i].filename);
                }
            });

            this.setState({alreadyCalled: true})
        });
    }

    render() {
        if (!this.state.alreadyCalled) {
            this.getList();
        }

        return (
          <div className="App-title">
              <h3> Welcome To RAudio. </h3>
              {<MusicList videos={this.state.list} />}
          </div>
        )
    }
}

export default App;
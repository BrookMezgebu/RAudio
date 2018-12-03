import React, {Component} from 'react'
import '../styles/App.css'
//import '../styles/materialize.min.css'
import '../styles/md1.css'
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
            alreadyCalled: false,
            isComplete: false
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
                    fetch(`http://localhost:8002/item_detail?file=${json[i].filepath}\\${json[i].filename}`)
                        .then(response => response.json())
                        .then(jsonx => {
                            this.addItem({
                                filename: json[i].filename,
                                filepath: json[i].filepath,
                                title: jsonx.title,
                                artist: jsonx.artist,
                                album: jsonx.album,
                                year: jsonx.year,
                                genre: jsonx.genre
                            });
                            console.log(this.state);
                            if (i === json.length - 1) {
                                this.setState({isComplete: true})
                            }
                        });
                }
            });

            this.setState({alreadyCalled: true})
        });
    }

    render() {
        if (!this.state.alreadyCalled) {
            this.getList();
        }

        if (!this.state.isComplete) {
            return (
                <div className="App-title">
                    <h3> Welcome To RAudio. </h3>
                </div>
            )
        } else {
            return (
                <div className="App-title">
                    <h3> Welcome To RAudio. </h3>
                    {<MusicList videos={this.state.list}/>}
                </div>
            )
        }
    }
}

export default App;
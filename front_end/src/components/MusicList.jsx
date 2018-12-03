import React , { Component } from 'react'
import Paper from "@material-ui/core/es/Paper/Paper";
import Grow from "@material-ui/core/es/Grow/Grow";

class MusicList extends Component {

    constructor(props) {
        super(props);
        this.state = {
            faded : false,
        }
    }

    componentDidMount() {
        setInterval(() => {this.setState({faded: true})} , 2000)
    }

    render() {
        const {videos} = this.props;
        console.log(videos);

        return (
            <div className="video-container_0 scrolling">
                {
                    videos.map(
                        (item , k) => {
                            return (
                                <Grow in={this.state.faded}>
                                    <Paper key={k} elevation={4} className="single-video card waves-effect waves-light">
                                        <img src={"./Test8.png"} width={'100%'} height={'100%'} className="video-image"/>
                                        <div className="video-detail">
                                            <span>{item.title}</span><br />
                                            <span>{item.artist}</span><br />
                                            <span>{item.genre}</span>
                                        </div>
                                    </Paper>
                                </Grow>
                            )
                        }
                    )
                }
            </div>
        )
    }

}

export default MusicList
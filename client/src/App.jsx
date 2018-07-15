import React, {Component} from 'react';
import { ToastContainer } from 'react-toastify';

import Messages from './Messages';

class App extends Component {
    render() {
        return (
            <div className="App">
                <header className="App-Header">
                    <h1 className="App-Title">All Messages</h1>
                </header>
                <div className="App-Intro">
                    <Messages />
                </div>
                <ToastContainer />
            </div>
        );
    }
}

export default App;
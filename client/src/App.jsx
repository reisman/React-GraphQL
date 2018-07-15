import React, {Component} from 'react';
import { Container, Row, Col } from 'reactstrap';

import Messages from './Messages';
import Users from './Users';
import AddMessage from './AddMessage';

class App extends Component {
    render() {
        return (
            <Container fluid>
                <Row>
                    <Col>
                        <h1 className="App-Title">React-GraphQL-Chat</h1>
                    </Col>
                </Row>
                <Row>
                    <Col xs="3"><Users /></Col>
                    <Col xs="9"><Messages /></Col>
                </Row>
                <Row>
                    <Col>
                        <AddMessage user={this.props.user}/>
                    </Col>
                </Row>
            </Container>
        );
    }
}

export default App;
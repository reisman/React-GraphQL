import React, { Component } from 'react';
import { InputGroup, Input, InputGroupAddon, Button } from 'reactstrap';
import { graphql } from 'react-apollo';
import gql from 'graphql-tag';

class AddMessage extends Component {
    constructor(props) {
        super(props);
        this.state = {
            text: '',
            mutate: props.mutate,
            user: props.user,
        };
    }
    
    render() {
        const handleChange = e => {
            e.preventDefault();
            this.setState({
                ...this.state,
                text: e.target.value,
            });
        };

        const handleKeyUp = e => {
            if (e.keyCode === 13) {
                e.preventDefault();
                postMessage.bind(this)();
            }
        };

        const handleClick = (e) => {
            e.preventDefault();
            postMessage.bind(this)();
        };

        const postMessage = () => {
            this.state.mutate({
                variables: {
                    text: this.state.text,
                    authorId:this.state.user,
                },
            }).then(_ => {
                this.setState({
                    ...this.state,
                    text: '',
                });
            });
        };

        return (
            <div>
                <br />
                <InputGroup>
                    <Input value={this.state.text} onChange={handleChange.bind(this)} onKeyUp={handleKeyUp.bind(this)} />
                    <InputGroupAddon addonType="append"><Button onClick={handleClick.bind(this)}>Send</Button></InputGroupAddon>
                </InputGroup>
            </div>
        );
    }
};

const addMessageMutation = gql`
  mutation addMessage($text: String!, $authorId: String!){
    addMessage(text: $text, authorId: $authorId) {
        id,
        text
    }
}
`;

export default graphql(addMessageMutation)(AddMessage);
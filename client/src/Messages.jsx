import React from 'react';
import { Query } from 'react-apollo';
import gql from 'graphql-tag';

import Message from './Message';

const GET_MESSAGES = gql`
{
    messages {
        id,
        text,
        publishedat,
        author {
            name
        }
    }
}
`

const Messages = () => {
    return (
        <Query query={GET_MESSAGES} pollInterval={1000}>
            {({ loading, error, data }) => {
                if (loading) return 'LOADING...';
                if (error) return `ERROR: ${error.message}`;
                return data.messages.map(msg => {
                    return (
                        <Message key={msg.id} text={msg.text} publishedat={msg.publishedat} author={msg.author.name} />
                    );
                });
            }}
        </Query>
    );
};

export default Messages;
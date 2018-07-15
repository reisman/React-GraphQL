import React, { Component } from 'react';
import { graphql, Query } from 'react-apollo';
import gql from 'graphql-tag';

const GET_MESSAGES = gql`
{
    messages {
        id,
        text
    }
}
`

const Messages = () => {
    return (
        <Query query={GET_MESSAGES}>
            {({ loading, error, data }) => {
                if (loading) return 'LOADING...';
                if (error) return `ERROR: ${error.message}`;
                return data.messages.map(msg => {
                    return (
                        <div key={msg.id}>
                            <p>{msg.text}</p>
                        </div>
                    );
                });
            }}
        </Query>
    );
};

export default Messages;
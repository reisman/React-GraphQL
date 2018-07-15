import React from 'react';
import { graphql, Query } from 'react-apollo';
import gql from 'graphql-tag';

import User from './User';

const GET_USERS = gql`
{
    authors {
        id,
        name
    }
}
`

const Users = (props) => {
    return (
        <Query query={GET_USERS} pollInterval={1000}>
            {({ loading, error, data }) => {
                if (loading) return 'LOADING...';
                if (error) return `ERROR: ${error.message}`;
                return data.authors.map(author => {
                    return (
                        <User key={author.id} name={author.name} />
                    );
                });
            }}
        </Query>
    );
};

export default Users;

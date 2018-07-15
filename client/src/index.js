import React from 'react';
import ReactDom from 'react-dom';
import { ApolloProvider } from 'react-apollo';
import { ApolloClient } from 'apollo-client';
import { HttpLink } from 'apollo-link-http';
import { InMemoryCache } from 'apollo-cache-inmemory';

import App from './App';
import 'bootstrap/dist/css/bootstrap.min.css';

const address = 'http://localhost:8080/graphql';
const client = new ApolloClient({
    link: new HttpLink({ uri: address }),
    cache: new InMemoryCache(),
});

ReactDom.render(
    <ApolloProvider client={client}>
        <App />
    </ApolloProvider>,
    document.getElementById('root')
);
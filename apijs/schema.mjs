import gql from 'graphql-tag';
import { GraphQLScalarType } from 'graphql';
import { getAuthors, getMessages, createAuthor, createMessage, findAuthor } from './dataaccess';

export const typeDefs = gql`
    scalar Date

    type Author {
        id: ID
        name: String
    }

    type Message {
        id: ID
        publishedat: Date
        text: String
        author: Author
    }

    type Query {
        authors: [Author]
        messages: [Message]
    }

    type Mutation {
        addMessage(text: String, authorId: String): Message
        addAuthor(name: String): Author
    }
`

export const resolvers = {
    Query: {
        authors: () => getAuthors(),
        messages: () => getMessages(),
    },
    Mutation: {
        addMessage: (_, {text, authorId}) => {
            const author = findAuthor(authorId);
            return createMessage(text, author);
        },
        addAuthor: function(_, { name }) {
            return createAuthor(name);
        },
    },
    Date: new GraphQLScalarType({
        name: 'Date',
        description: 'Date custom scalar type',
        parseValue(value) {
          return new Date(value);
        },
        serialize(value) {
          return value.getTime();
        },
        parseLiteral(ast) {
          if (ast.kind === Kind.INT) {
            return new Date(ast.value);
          }
          return null;
        },
    }),
}
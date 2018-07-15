import ApolloServer from 'apollo-server';
import gql from 'graphql-tag';
import uuid from 'uuid';
import { GraphQLScalarType } from 'graphql';
import moment from 'moment';

const createId = () => {
    return uuid.v4();
}

const sample_authors = [
    {id: "1", name: "Hans"},
    {id: "2", name: "Peter"},
    {id: "3", name: "Sam"},
    {id: "4", name: "Tom"},
    {id: "5", name: "Joe"},
];

const sample_messages = [
    {id: "1", author: sample_authors[4], publishedat: moment().subtract(1, 'seconds').toDate(), text: "TextA"},
    {id: "2", author: sample_authors[2], publishedat: moment().subtract(2, 'seconds').toDate(), text: "TextB"},
    {id: "3", author: sample_authors[2], publishedat: moment().subtract(3, 'seconds').toDate(), text: "TextC"},
    {id: "4", author: sample_authors[3], publishedat: moment().subtract(4, 'seconds').toDate(), text: "TextE"},
    {id: "5", author: sample_authors[4], publishedat: moment().subtract(5, 'seconds').toDate(), text: "TextF"},
];

const typeDefs = gql`
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

const resolvers = {
    Query: {
        authors: () => sample_authors,
        messages: () => sample_messages,
    },
    Mutation: {
        addMessage: (_, {text, authorId}) => {
            const id = createId();
            const publishedat = moment().toDate();
            const author = sample_authors.find(a => a.id === authorId);
            const msg = {id, author, publishedat, text};
            sample_messages.push(msg);
            return msg;
        },
        addAuthor: function(_, { name }) {
            const id = createId();
            const author = {id, name};
            sample_authors.push(author);
            return author;
        },
    },
    Date: new GraphQLScalarType({
        name: 'Date',
        description: 'Date custom scalar type',
        parseValue(value) {
          return new Date(value); // value from the client
        },
        serialize(value) {
          return value.getTime(); // value sent to the client
        },
        parseLiteral(ast) {
          if (ast.kind === Kind.INT) {
            return new Date(ast.value) // ast value is always in string format
          }
          return null;
        },
      }),
}

const server = new ApolloServer.ApolloServer({
    typeDefs,
    resolvers,
    cors: true,
});

server.listen({ port: 8080 }).then(({ url }) => {
    console.log(`Server listening at ${url}`);
});
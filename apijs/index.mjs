import ApolloServer from 'apollo-server';
import { typeDefs, resolvers } from './schema';

const server = new ApolloServer.ApolloServer({
    typeDefs,
    resolvers,
    cors: true,
});

server.listen({ port: 8080 }).then(({ url }) => {
    console.log(`Server listening at ${url}`);
});
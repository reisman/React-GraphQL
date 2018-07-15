import chance from 'chance';
import gql from 'graphql-tag';

const addAuthorMutation = gql`
  mutation addAuthor($name: String!){
    addAuthor(name: $name) {
        id
    }
}
`;

export const initLoginUser = async client => {
    const userName = chance().name();
    return client.mutate({
        mutation: addAuthorMutation,
        variables: {
            name: userName,
        },
    }).then(res => res.data.addAuthor.id);
};
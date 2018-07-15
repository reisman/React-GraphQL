package data

import (
	"errors"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/relay"
	"golang.org/x/net/context"
)

var Schema *graphql.Schema

func init() {
	var authorType *graphql.Object
	var messageType *graphql.Object

	nodeDefinitions := createNodeDefinitions(authorType, messageType)
	authorType = createAuthorType(nodeDefinitions)
	messageType = createMessageType(nodeDefinitions, authorType)
	queryType := createQuery(nodeDefinitions, authorType, messageType)
	mutationType := createMutation(authorType, messageType)

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    queryType,
		Mutation: mutationType,
	})

	if err != nil {
		panic(err)
	}

	Schema = &schema
}

func createNodeDefinitions(authorType *graphql.Object, messageType *graphql.Object) *relay.NodeDefinitions {
	return relay.NewNodeDefinitions(relay.NodeDefinitionsConfig{
		IDFetcher: func(id string, info graphql.ResolveInfo, ctx context.Context) (interface{}, error) {
			resolvedID := relay.FromGlobalID(id)
			switch resolvedID.Type {
			case "Author":
				return getAuthor(resolvedID.ID), nil
			case "Message":
				return getMessage(resolvedID.ID), nil
			default:
				return nil, errors.New("Unknown type: " + resolvedID.Type)
			}
		},
		TypeResolve: func(p graphql.ResolveTypeParams) *graphql.Object {
			switch p.Value.(type) {
			case *Author:
				return authorType
			default:
				return messageType
			}
		},
	})
}

func createQuery(nodeDefinitions *relay.NodeDefinitions, authorType *graphql.Object, messageType *graphql.Object) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"authors": &graphql.Field{
				Type: graphql.NewList(authorType),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return getAuthors(), nil
				},
			},
			"messages": &graphql.Field{
				Type: graphql.NewList(messageType),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return getMessages(), nil
				},
			},
			"node": nodeDefinitions.NodeField,
		},
	})
}

func createMutation(authorType *graphql.Object, messageType *graphql.Object) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"addMessage": &graphql.Field{
				Type: messageType,
				Args: graphql.FieldConfigArgument{
					"text": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"authorId": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					text, _ := p.Args["text"].(string)
					authorID, _ := p.Args["authorId"].(string)
					author := getAuthor(authorID)
					msg := createMessage(text, author)
					addMessage(msg)
					return msg, nil
				},
			},
			"addAuthor": &graphql.Field{
				Type: authorType,
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					name, _ := p.Args["name"].(string)
					author := createAuthor(name)
					addAuthor(author)
					return author, nil
				},
			},
		},
	})
}

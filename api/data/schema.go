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

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: queryType,
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

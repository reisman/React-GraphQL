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

	nodeDefinitions := relay.NewNodeDefinitions(relay.NodeDefinitionsConfig{
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

	authorType = graphql.NewObject(graphql.ObjectConfig{
		Name:        "Author",
		Description: "The author of a published message",
		Fields: graphql.Fields{
			"id": relay.GlobalIDField("Author", nil),
			"name": &graphql.Field{
				Type:        graphql.String,
				Description: "The name of the author",
			},
		},
		Interfaces: []*graphql.Interface{
			nodeDefinitions.NodeInterface,
		},
	})

	messageType = graphql.NewObject(graphql.ObjectConfig{
		Name:        "Message",
		Description: "A message that was published by an author",
		Fields: graphql.Fields{
			"id": relay.GlobalIDField("Message", nil),
			"text": &graphql.Field{
				Type:        graphql.String,
				Description: "The content of the message",
			},
			"publishedat": &graphql.Field{
				Type:        graphql.DateTime,
				Description: "The time when the message was published",
			},
			"author": &graphql.Field{
				Type:        authorType,
				Description: "The author of this message",
			},
		},
		Interfaces: []*graphql.Interface{
			nodeDefinitions.NodeInterface,
		},
	})

	queryType := graphql.NewObject(graphql.ObjectConfig{
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

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: queryType,
	})

	if err != nil {
		panic(err)
	}

	Schema = &schema
}

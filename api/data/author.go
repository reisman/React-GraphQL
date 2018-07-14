package data

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/relay"
)

type Author struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func createAuthorType(nodeDefinitions *relay.NodeDefinitions) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
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
}

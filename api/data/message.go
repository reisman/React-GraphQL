package data

import (
	"time"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/relay"
)

type Message struct {
	ID          string    `json:"id"`
	Text        string    `json:"text"`
	PublishedAt time.Time `json:"publishedat"`
	Author      *Author   `json:"author"`
}

func createMessage(text string, author *Author) *Message {
	return &Message{
		ID:          generateID(),
		Author:      author,
		PublishedAt: time.Now(),
		Text:        text,
	}
}

func createMessageType(nodeDefinitions *relay.NodeDefinitions, authorType *graphql.Object) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
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
}

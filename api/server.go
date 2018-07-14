package main

import (
	"GraphQLMessages/api/data"
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

func main() {
	log.Print("Crateing graphql schema")
	schema := *data.Schema

	hdl := createHandler(schema)
	http.Handle("/graphql", hdl)

	port := ":8080"
	log.Printf(`Starting GraphQL server on http://localhost%v`, port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalf("Starting server failed: %v", err)
	}
}

func createHandler(schema graphql.Schema) *handler.Handler {
	return handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})
}

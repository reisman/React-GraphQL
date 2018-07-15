package main

import (
	"GraphQLMessages/api/data"
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/rs/cors"
)

func main() {
	log.Print("Creating graphql schema")
	schema := *data.Schema

	mux := http.NewServeMux()
	graphqlHandler := createHandler(schema)
	mux.Handle("/graphql", graphqlHandler)
	corsHandler := cors.Default().Handler(mux)

	port := ":8080"
	log.Printf(`Starting GraphQL server on http://localhost%v`, port)
	err := http.ListenAndServe(port, corsHandler)
	if err != nil {
		log.Fatalf("Starting server failed: %v", err)
	}
}

func createHandler(schema graphql.Schema) *handler.Handler {
	return handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: false,
	})
}

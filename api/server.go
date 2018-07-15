package main

import (
	"GraphQLMessages/api/data"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/http2"

	"github.com/gorilla/handlers"
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
	handler := withCompression(withLogging(withCors(mux)))
	
	server := http.Server{
		Addr:    ":8080",
		Handler: handler,
	}
	
	log.Printf(`Starting GraphQL server on http://localhost%v`, server.Addr)
	
	http2.ConfigureServer(&server, &http2.Server{})
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("Starting server failed: %v", err)
	}
}

func withCors(handler http.Handler) http.Handler {
	return cors.Default().Handler(handler)
}

func withCompression(handler http.Handler) http.Handler {
	return handlers.CompressHandler(handler)
}

func withLogging(handler http.Handler) http.Handler {
	return handlers.LoggingHandler(os.Stdout, handler)
}

func createHandler(schema graphql.Schema) *handler.Handler {
	return handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: false,
	})
}

package main

import (
	"log"
	"net/http"

	"github.com/graphql-go/handler"
	"github.com/graphql-go/graphql"
)

func main() {
	log.Print("Crateing graphql schema")
	schema, err := createSchema()
	if err != nil {
		log.Fatalf("Failed to create graphql schema: %v", err)
	}

	hdl := createHandler(schema)
	http.Handle("/graphql", hdl)

	port := ":8080"
	log.Printf(`Starting GraphQL server on http://localhost%v`, port)
	err = http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalf("Starting server failed: %v", err)
	}

	/*
		if err != nil {
			log.Fatalf("Failed to create new graphql schema: %v", err)
		}

		query := `
			{
				hello
			}
		`

		params := graphql.Params{Schema: schema, RequestString: query}
		r := graphql.Do(params)
		if r.HasErrors() {
			log.Fatalf("Failed to execute graphql query: %+v", r.Errors)
		}

		result, _ := json.Marshal(r)
		fmt.Printf("%s \n", result)*/
}

func createHandler(schema graphql.Schema) *handler.Handler {
	return handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})
}

func createSchema() (graphql.Schema, error) {
	fields := graphql.Fields{
		"hello": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "world", nil
			},
		},
	}

	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	return graphql.NewSchema(schemaConfig)
}

package main

import (
	"io/ioutil"
	"log"
	"net/http"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

func main() {
	// Get Schema,
	s, err := getSchema("./schema.gql")
	if err != nil {
		log.Println(err)
	} // and parse it.
	schema := graphql.MustParseSchema(s, &Resolver{})

	// Handle http graph-queries based on schema.
	http.Handle("/api", &relay.Handler{Schema: schema})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// If there is an error return an empty schema.
func getSchema(path string) (string, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

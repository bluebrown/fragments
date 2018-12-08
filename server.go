package main

import (
	"io/ioutil"
	"log"
	"net/http"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

func main() {

	store.init() // Init the store with some mock data.

	// Read schema,
	s, err := getSchema("./schema.gql")
	if err != nil {
		log.Println("No shema found!", err)
	} // and parse it.
	schema := graphql.MustParseSchema(s, &Resolver{})

	// Handle http graph-queries based on schema.
	http.Handle("/api", &relay.Handler{Schema: schema})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

/* getSchema reads file from specified path into
   byte-buffer. Returns an empty byte if an
   error occurs, for revovery purposes. */
func getSchema(path string) (string, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

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
	// Spin up static File Server
	http.Handle("/", fragmentServer())
	// Read, parse and serve schema.
	http.Handle("/api", schemaHandler("./schema.gql"))
	// Serve GraphQL API
	log.Fatalln(http.ListenAndServe(":8080", nil))
}

// schemaHandler reads and parses a schema file.
// It returns a handler that can get used by an http router.
func schemaHandler(path string) *relay.Handler {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalln(err)
	}
	return &relay.Handler{ // ! Here is where the magic happens!
		Schema: graphql.MustParseSchema(string(b), &Resolver{}),
	}
}

// fragmentserver is a static fileserver with
// basic configuration to serve the service fragment.
func fragmentServer() http.Handler {
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	return fs
}

package server

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/bluebrown/fragments/resolver"
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

// Config holds the configuration for the server.
type Config struct {
	StaticPattern string
	APIPattern    string
	SchemaPath    string
	Port          string
}

// Run takes configuration arguments, reads&parses the
// graphql-schema and spins up the server with static directory.
func Run(conf *Config) error {
	// Spin up static File Server
	http.Handle(conf.StaticPattern, FragmentServer())
	// Read, parse and serve schema.
	http.Handle(conf.APIPattern, SchemaHandler(conf.SchemaPath))
	// Serve GraphQL API
	return (http.ListenAndServe(conf.Port, nil))
}

// SchemaHandler reads and parses a schema file.
// It returns a handler that can get used by an http router.
func SchemaHandler(path string) *relay.Handler {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalln(err)
	}
	return &relay.Handler{ // ! Here is where the magic happens!
		Schema: graphql.MustParseSchema(string(b), &resolver.Resolver{}),
	}
}

// FragmentServer is a static fileserver with
// basic configuration to serve the service fragment.
func FragmentServer() http.Handler {
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	return fs
}

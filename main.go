package main

import (
	"log"

	"github.com/bluebrown/fragments/server"
)

// SConfig referrs to server-configuration.
var SConfig = &server.Config{
	StaticPattern: "/",
	APIPattern:    "/api",
	SchemaPath:    "./schema.gql",
	Port:          ":8080",
}

func main() {

	log.Println("Running on port" + SConfig.Port)
	log.Fatalln(server.Run(SConfig))
}

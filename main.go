package main

import (
	"log"

	"github.com/bluebrown/fragments/model"
	"github.com/bluebrown/fragments/server"
)

// SConfig referrs to server-configuration.
var SConfig = &server.Config{
	StaticPattern: "/",
	APIPattern:    "/api",
	SchemaPath:    "./gtype/schema.gql",
	Port:          ":8080",
}

func main() {

	model.Example()

	log.Println("Running on port" + SConfig.Port)
	log.Fatalln(server.Run(SConfig))
}

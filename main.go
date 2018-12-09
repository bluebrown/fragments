package main

// Model Singleton
var Model = driver{
	model: model{},
}

func main() {
	store.init() // Init the store with some mock data.

	// Start Server
	newServer(serverConf{
		staticPattern: "/",
		apiPattern:    "/api",
		schemaPath:    "./schema.gql",
		port:          ":8080",
	})
}

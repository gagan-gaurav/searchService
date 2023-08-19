package main

import (
	"fmt"
	"net/http"
	"search/internal/es"
	"search/internal/routes"
)

func main() {
	routes.SetRouter()                 // Set the routes.
	es.CreateClient()                  // Create Elasticsearch client
	es.CreateProjectsIndex("projects") // Create 'projects' index if already not present.

	// Start the server at port 8081
	port := "8081"
	fmt.Printf("Server started on port %s\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}

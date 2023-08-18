package main

import (
	"fmt"
	"net/http"
	"search/internal/es"
	"search/internal/routes"
)

func main() {
	routes.SetRouter()
	es.CreateClient()

	// Start the server at port 8081
	port := "8081"
	fmt.Printf("Server started on port %s\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}

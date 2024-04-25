package main

/*

Package main is the entry point for the application.

Author: Ibrahim Alaybeyi

Description:
This application demonstrates how to create a simple web server in Go
using the standard net/http package with a modular routing approach.

The main package defines routes for various operations and registers
them with their corresponding handlers. Each handler is responsible
for handling HTTP requests related to a specific domain or resource.

*/

import (
	"g-route/controller"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	controller.RegisterRoutes(mux)
	http.ListenAndServe("localhost:8080", mux)
}

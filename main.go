package main

import (
	"go-echo-server/routes"
	"log"
	"net/http"

	errorsHandler "github.com/tacheraSasi/tripwire/errorsHandler"
)
const PORT = ":3000"
func main() {
	mux := http.NewServeMux()
	routes.HandleRoutes(mux)
	log.Println("Server is running on port",PORT)
	error := http.ListenAndServe(PORT, mux)
	errorsHandler.Check(error, "Failed to run the server")
}

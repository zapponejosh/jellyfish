package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zapponejosh/jellyApi/dbOps"
)

func main() {
	r := mux.NewRouter()

	// API routes
	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/hello", handlerApi)

	// Start server
	fmt.Println("Starting the server on http://localhost:3001")
	log.Fatal(http.ListenAndServe(":3001", r))

}

func handlerApi(w http.ResponseWriter, r *http.Request) {
	str := dbOps.TestQuery()

	json.NewEncoder(w).Encode(map[string]string{"message": "Hello from the Go API -- Check this out for serving the react app too: https://github.com/gorilla/mux#serving-single-page-applications", "dataTest": str})
}

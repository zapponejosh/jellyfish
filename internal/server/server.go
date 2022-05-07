package server

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/zapponejosh/jellyfish/internal/dbops"
	"github.com/zapponejosh/jellyfish/internal/server/handlers"
	"github.com/zapponejosh/jellyfish/internal/settings"
)

func New(appSettings *settings.Settings, db *dbops.DB) (*http.Server, error) {

	r := mux.NewRouter()

	// API routes
	api := r.PathPrefix("/api").Subrouter()
	api.Handle("/hello", handlers.NewGetHandler(db))

	// Start server
	// fmt.Println("Starting the server on http://localhost:3001")
	// log.Fatal(http.ListenAndServe(":3001", r))

	// r.Handle("/user/", users.NewListHandler(db)).Methods(http.MethodGet)
	// r.Handle("/user/", users.NewCreateHandler(db)).Methods(http.MethodPost)

	// r.Handle("/user/{id:[0-9]+}/", users.NewGetHandler(db)).Methods(http.MethodGet)
	// r.Handle("/user/{id:[0-9]+}/", users.NewUpdateHandler(db)).Methods(http.MethodPut)
	// r.Handle("/user/{id:[0-9]+}/", users.NewDeleteHandler(db)).Methods(http.MethodDelete)

	return &http.Server{Handler: r, Addr: appSettings.ServerAddress}, nil
}

// func handlerApi(w http.ResponseWriter, r *http.Request, db *dbops.DB) {
// 	str := dbops.TestQuery()

// 	json.NewEncoder(w).Encode(map[string]string{"message": "Hello from the Go API -- Check this out for serving the react app too: https://github.com/gorilla/mux#serving-single-page-applications", "dataTest": str})
// }

package server

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/zapponejosh/jellyfish/internal/dbops"
	"github.com/zapponejosh/jellyfish/internal/server/handlers"
	"github.com/zapponejosh/jellyfish/internal/server/handlers/users"
	"github.com/zapponejosh/jellyfish/internal/settings"
)

func New(appSettings *settings.Settings, db *dbops.DB) (*http.Server, error) {

	r := mux.NewRouter()

	// API routes
	api := r.PathPrefix("/api").Subrouter()
	api.Handle("/hello", handlers.NewGetHandler(db))
	api.Handle("/user", users.NewCreateHandler(db)).Methods(http.MethodPost)
	api.Handle("/user/{id:[0-9]+}", users.NewGetHandler(db)).Methods(http.MethodGet)

	return &http.Server{Handler: r, Addr: appSettings.ServerAddress}, nil
}

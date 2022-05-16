package handlers

import (
	"context"
	"errors"
	"log"
	"net/http"

	// "github.com/gorilla/mux"
	"github.com/jackc/pgx/v4"
	"github.com/zapponejosh/jellyfish/internal/models"
	"github.com/zapponejosh/jellyfish/internal/server/handlers/responders"
)

type TestGetter interface {
	GetTest(ctx context.Context) (*models.Test, error)
}

func NewGetHandler(getter TestGetter) *GetHandler {
	return &GetHandler{getter: getter}
}

type GetHandler struct {
	getter TestGetter
}

func (h GetHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	testR, err := h.getter.GetTest(r.Context())
	if errors.Is(err, pgx.ErrNoRows) {
		responders.Error(w, "not found", http.StatusNotFound)
		return
	}
	if err != nil {
		log.Println("error:", err)
		responders.Error(w, "error", http.StatusInternalServerError)
		return
	}
	res := map[string]string{"message": "Hello from the Go API -- Check this out for serving the react app too: https://github.com/gorilla/mux#serving-single-page-applications", "dataTest": testR.TestStr}

	responders.OK(w, res)
}

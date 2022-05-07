package handlers

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"net/http"

	// "github.com/gorilla/mux"
	"github.com/zapponejosh/jellyfish/internal/server/handlers/responders"
)

type Test struct {
	testStr string
}

type TestGetter interface {
	GetTest(ctx context.Context) (*Test, error)
}

func NewGetHandler(getter TestGetter) *GetHandler {
	return &GetHandler{getter: getter}
}

type GetHandler struct {
	getter TestGetter
}

func (h GetHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user, err := h.getter.GetTest(r.Context())
	if errors.Is(err, sql.ErrNoRows) {
		responders.Error(w, "user not found", http.StatusNotFound)
		return
	}
	if err != nil {
		log.Println("error getting user:", err)
		responders.Error(w, "error getting user", http.StatusInternalServerError)
		return
	}

	responders.OK(w, user)
}

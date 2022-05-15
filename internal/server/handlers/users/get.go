package users

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/zapponejosh/jellyfish/internal/models"
	"github.com/zapponejosh/jellyfish/internal/server/handlers/responders"
)

type UserGetter interface {
	GetUser(ctx context.Context, id int) (*models.User, error)
}

func NewGetHandler(getter UserGetter) *GetHandler {
	return &GetHandler{getter: getter}
}

type GetHandler struct {
	getter UserGetter
}

func (h GetHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Println("error parsing id:", err)
		responders.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	user, err := h.getter.GetUser(r.Context(), id)
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

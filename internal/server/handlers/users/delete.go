package users

import (
	"context"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4"

	"github.com/zapponejosh/jellyfish/internal/server/handlers/responders"
)

type UserDeleter interface {
	DeleteUser(ctx context.Context, id int) (int, error)
}

func NewDeleteHandler(deleter UserDeleter) *DeleteHandler {
	return &DeleteHandler{deleter: deleter}
}

type DeleteHandler struct {
	deleter UserDeleter
}

func (h DeleteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Println("error parsing id:", err)
		responders.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	_, err = h.deleter.DeleteUser(r.Context(), id)
	if errors.Is(err, pgx.ErrNoRows) {
		responders.Error(w, "user not found", http.StatusNotFound)
		return
	}
	if err != nil {
		log.Println("error deleting user:", err)
		responders.Error(w, "error deleting user", http.StatusInternalServerError)
		return
	}

	responders.NoContent(w)
}

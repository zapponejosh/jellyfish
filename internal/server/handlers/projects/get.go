package projects

import (
	"context"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4"
	"github.com/zapponejosh/jellyfish/internal/models"
	"github.com/zapponejosh/jellyfish/internal/server/handlers/responders"
)

type ProjectGetter interface {
	GetProject(ctx context.Context, id int) (*models.GetProject, error)
}

type GetHandler struct {
	getter ProjectGetter
}

func NewGetHandler(getter ProjectGetter) *GetHandler {
	return &GetHandler{getter: getter}
}

func (h GetHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Println("error parsing id:", err)
		responders.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	project, err := h.getter.GetProject(r.Context(), id)

	if errors.Is(err, pgx.ErrNoRows) {
		responders.Error(w, "no project not found", http.StatusNotFound)
		return
	}

	if err != nil {
		log.Println("error getting project:", err)
		responders.Error(w, "error getting project", http.StatusInternalServerError)
		return
	}

	responders.OK(w, project)
}

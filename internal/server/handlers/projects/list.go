package projects

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/jackc/pgx/v4"
	"github.com/zapponejosh/jellyfish/internal/models"
	"github.com/zapponejosh/jellyfish/internal/server/handlers/responders"
)

type ProjectLister interface {
	ListProjects(ctx context.Context) ([]*models.ProjectPreview, error)
}

type ListHandler struct {
	lister ProjectLister
}

func NewListHandler(lister ProjectLister) *ListHandler {
	return &ListHandler{lister: lister}
}

func (h ListHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	projects, err := h.lister.ListProjects(r.Context())

	if errors.Is(err, pgx.ErrNoRows) {
		responders.Error(w, "no projects not found", http.StatusNotFound)
		return
	}

	if err != nil {
		log.Println("error getting projects:", err)
		responders.Error(w, "error getting projects", http.StatusInternalServerError)
		return
	}

	responders.OK(w, projects)
}

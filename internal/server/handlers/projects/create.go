package projects

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/jackc/pgx/v4"
	"github.com/zapponejosh/jellyfish/internal/models"
	"github.com/zapponejosh/jellyfish/internal/server/handlers/responders"
)

type ProjectCreator interface {
	CreateProject(ctx context.Context, project *models.Project) (int, error)
}

type CreateHandler struct {
	creator ProjectCreator
}

func NewCreateHandler(creator ProjectCreator) *CreateHandler {
	return &CreateHandler{creator: creator}
}

func (h CreateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var project models.Project
	if err := json.NewDecoder(r.Body).Decode(&project); err != nil {
		log.Println("error decoding body:", err)
		responders.Error(w, "invalid project data", http.StatusBadRequest)
		return
	}
	projectID, err := h.creator.CreateProject(r.Context(), &project)

	if errors.Is(err, pgx.ErrNoRows) {
		responders.Error(w, "Error returning id", http.StatusNotFound)
		return
	}

	if err != nil {
		log.Println("error creating project:", err)
		responders.Error(w, "error creating project", http.StatusInternalServerError)
		return
	}

	responders.OK(w, projectID)
}

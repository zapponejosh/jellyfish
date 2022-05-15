package users

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/zapponejosh/jellyfish/internal/models"
	"github.com/zapponejosh/jellyfish/internal/server/handlers/responders"
)

type UserCreator interface {
	CreateUser(context.Context, *models.User) (int, error)
}

func NewCreateHandler(creator UserCreator) *CreateHandler {
	return &CreateHandler{creator: creator}
}

type CreateHandler struct {
	creator UserCreator
}

func (h CreateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Println("error decoding body:", err)
		responders.Error(w, "invalid user data", http.StatusBadRequest)
		return
	}
	id, err := h.creator.CreateUser(r.Context(), &user)
	if err != nil {
		log.Println("error creating user:", err)
		responders.Error(w, "error creating user", http.StatusInternalServerError)
		return
	}
	responders.Created(w, id)
}

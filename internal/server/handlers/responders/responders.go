package responders

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func Respond(w http.ResponseWriter, msg interface{}, code int) {
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(msg); err != nil {
		log.Println("error writing message")
	}
}

func Error(w http.ResponseWriter, msg string, code int) {
	Respond(w, map[string]string{"error": msg}, code)
}

func Created(w http.ResponseWriter, id int) {
	w.Header().Add("Location", strconv.Itoa(id))
	w.WriteHeader(http.StatusCreated)
}

func OK(w http.ResponseWriter, obj interface{}) {
	if obj == nil {
		return
	}
	if err := json.NewEncoder(w).Encode(obj); err != nil {
		log.Println("error encoding object:", err)
	}
}

func NoContent(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}

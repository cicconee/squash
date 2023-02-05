package team

import (
	"encoding/json"
	"net/http"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var resp = Payload{
		"message": "Your team was created!",
		"id":      "asdfafdasdf",
		"name":    "ywots",
	}

	JSONResponse(w, http.StatusOK, resp)
}

type Payload map[string]interface{}

func JSONResponse(w http.ResponseWriter, code int, payload Payload) {
	data, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}

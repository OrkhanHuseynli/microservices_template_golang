package handlers

import (
	"context"
	"encoding/json"
	"github.com/lovoo/goka"
	"microservices_template_golang/payment/src/models"
	"net/http"
)

type ServiceHandler struct {
}

func NewServiceHandler() http.Handler {
	return ServiceHandler{}
}

func (h ServiceHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost{
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var req models.Message
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	if req.Content == "" {
		http.Error(w, "Required body fields are empty", http.StatusUnprocessableEntity)
		return
	}

	c := context.WithValue(r.Context(), "content", req.Content)
	r = r.WithContext(c)



	//emitMessage()
	response := models.SimpleResponse{Message: "Your message was successfully put in the process" }
	encoder := json.NewEncoder(w)
	encoder.Encode(&response)
}


func emitMessage(emitter *goka.Emitter, m models.Message) error {
		return emitter.EmitSync(m.To, &m)
}

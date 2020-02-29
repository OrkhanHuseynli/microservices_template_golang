package handlers

import (
	"encoding/json"
	"github.com/lovoo/goka"
	"log"
	"net/http"
	"microservices_template_golang/payment/src/models"
)

type ServiceHandler struct {
	emitter *goka.Emitter
}

func NewServiceHandler(emitter *goka.Emitter) http.Handler {
	return ServiceHandler{emitter}
}

func (h ServiceHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost{
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var payment models.Payment
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&payment)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	if payment.Author == "" ||  payment.Sum == "" || payment.Product == "" {
		http.Error(w, "Required body fields are empty", http.StatusUnprocessableEntity)
		return
	}

	err = h.emitter.EmitSync(payment.Author, payment)
	if err != nil {
		log.Fatalf("error emitting message: %v", err)
		http.Error(w, "Error when trying to process the payment", http.StatusInternalServerError)
		return
	}
	response := models.SimpleResponse{Message: "Your payment was successfully put in the process" }
	encoder := json.NewEncoder(w)
	encoder.Encode(&response)
}

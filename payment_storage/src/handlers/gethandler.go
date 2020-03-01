package handlers

import (
	"encoding/json"
	"microservices_template_golang/payment/src/models"
	"net/http"
)

type GetHandler struct {
}

func NewGetHandler() http.Handler {
	return GetHandler{}
}

func (h GetHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		h.handleGetRequest(w, r)
	default :
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func (h GetHandler) handleGetRequest(w http.ResponseWriter, r *http.Request){
	paymentID := r.URL.Path[len("/database/"):]
	msg := "List of requested payments for " + paymentID
	response := models.PaymentResponse{Message: msg, Payments: nil}
	encoder := json.NewEncoder(w)
	encoder.Encode(&response)
}

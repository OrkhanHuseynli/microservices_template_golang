package handlers

import (
	"encoding/json"
	"fmt"
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
	paymentID := r.URL.Path[len("/payment/"):]
	fmt.Println(paymentID)
	resp, err := http.Get("http://localhost:8081/database/"+paymentID)
	if err != nil {
		fmt.Errorf(err.Error())
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	encoder := json.NewEncoder(w)
	encoder.Encode(resp)
}

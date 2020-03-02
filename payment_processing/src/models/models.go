package models

type SimpleResponse struct {
	Message string `json:"message"`
	Date    string `json:"date, omitempty"`
}

type Payment struct {
	Author    string	`json:"author"`
	Product   string	`json:"product"`
	Sum		  string	`json:"sum"`
}

type ProcessedPayment struct {
	PaymentID string 	`json:"paymentID"`
	Payment
}
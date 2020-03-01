package models

type SimpleResponse struct {
	Message string `json:"message"`
	Date    string `json:"date, omitempty"`
}

type PaymentResponse struct {
	Message string `json:"message"`
	Payments []Payment `json:"payments"`
}


type Payment struct {
	Author    string	`json:"author"`
	Product   string	`json:"product"`
	Sum		  string	`json:"sum"`
}
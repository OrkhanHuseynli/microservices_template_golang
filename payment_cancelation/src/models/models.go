package models

type Message struct {
	From    string
	To      string
	Content string
}

type SimpleResponse struct {
	Message string `json:"message"`
	Date    string `json:"date, omitempty"`
}

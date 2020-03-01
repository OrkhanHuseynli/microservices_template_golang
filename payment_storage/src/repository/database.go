package repository

import (
	"log"
	"microservices_template_golang/payment_storage/src/models"
)


type Database struct {
	db map[string]models.ProcessedPayment
}

func InitDatabase() *Database {
	return &Database{make(map[string]models.ProcessedPayment)}
}

func (d *Database)Store(payment *models.ProcessedPayment)  {
	d.db[payment.PaymentID] = *payment
	log.Printf("Database: payment %v was successfuly stored", payment.PaymentID)
	log.Printf("Total number of payments processed: %v", len(d.db))
}

func (d *Database) Get(paymentID string) models.ProcessedPayment  {
	return d.db[paymentID]
}
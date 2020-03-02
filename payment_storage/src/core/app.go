package core

import (
	"github.com/lovoo/goka"
	"log"
	"microservices_template_golang/payment_storage/src/eventmanager"
	"microservices_template_golang/payment_storage/src/models"
	"microservices_template_golang/payment_storage/src/repository"
)


var (
	brokers             = []string{"kafka:9090"}
	topic   goka.Stream = "payment-storage"
	group   goka.Group  = "example-group"
)

type App struct {

}

func New() *App {
	return &App{}
}

func (a *App) Start() {

	db := repository.InitDatabase()
	p := eventmanager.EventProcessor{}
	cb := func(ctx goka.Context, msg interface{}) {
		log.Printf("msg = %v", msg)
		payment, ok := msg.(*models.ProcessedPayment)
		if !ok {
			log.Println("Error while parsing message to the structure")
		}
		log.Printf("Payment with %v ID from %v was just processed", payment.PaymentID, payment.Author)
		db.Store(payment)
	}
	p.InitDefaultProcessor(brokers, group, topic, cb)
	p.Run()
}

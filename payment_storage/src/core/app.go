package core

import (
	"fmt"
	"github.com/lovoo/goka"
	"log"
	//"microservices_template_golang/payment_storage/src/eventmanager"
	"microservices_template_golang/payment_storage/src/handlers"
	//"microservices_template_golang/payment_storage/src/models"
	//"microservices_template_golang/payment_storage/src/repository"
	"net/http"
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

	//db := repository.InitDatabase()
	//p := eventmanager.EventProcessor{}
	//cb := func(ctx goka.Context, msg interface{}) {
	//	log.Printf("msg = %v", msg)
	//	payment, ok := msg.(*models.ProcessedPayment)
	//	if !ok {
	//		log.Println("Error while parsing message to the structure")
	//	}
	//	log.Printf("Payment with %v ID from %v was just processed", payment.PaymentID, payment.Author)
	//	db.Store(payment)
	//}

	port := 8081
	getHandler := handlers.NewGetHandler()
	http.Handle("/database/", getHandler)
	log.Printf("Server starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))

	//p.InitDefaultProcessor(brokers, group, topic, cb)
	//p.Run()
}

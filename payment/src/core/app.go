package core

import (
	"fmt"
	"github.com/lovoo/goka"
	"log"
	"microservices_template_golang/payment/src/eventmanager"
	"microservices_template_golang/payment/src/handlers"
	"microservices_template_golang/payment/src/models"
	"net/http"
)


var (
	brokers             = []string{"kafka:9090"}
	topic   goka.Stream = "example-stream"
	group   goka.Group  = "example-group"
)

type App struct {

}

func New() *App {
	return &App{}
}

func (a *App) Start() {
	port := 8080
	emitter := eventmanager.DefaultEmitter(brokers, topic, new(models.PaymentCodec))
	handler := handlers.NewServiceHandler(emitter)
	http.Handle("/payment", handler)
	log.Printf("Server starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

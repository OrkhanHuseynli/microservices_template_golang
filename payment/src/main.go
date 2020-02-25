package main

import (
	"github.com/lovoo/goka"
	"github.com/lovoo/goka/codec"
	"log"
	"microservices_template_golang/payment/src/core"
	"microservices_template_golang/payment/src/eventmanager"
)

var (
	brokers             = []string{"kafka:9090"}
	topic   goka.Stream = "example-stream"
	group   goka.Group  = "example-group"
)

// emits a single message and leave
func runEmitter() {
	pc := eventmanager.NewConfig()
	emitter := eventmanager.NewAppEmitter(brokers, topic, new(codec.String), pc)
	//defer emitter.Finish()
	err := emitter.EmitSync("some-key", "some-value")
	if err != nil {
		log.Fatalf("error emitting message: %v", err)
	}
	log.Println("******** message was successfully emitted ********")
}

func main() {
	runEmitter() // emits one message and stops
	app := core.New()
	app.Start()
}

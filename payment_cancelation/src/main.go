package main

import (
	"context"
	"fmt"
	"github.com/lovoo/goka"
	"github.com/lovoo/goka/codec"
	"github.com/lovoo/goka/kafka"
	"log"
	"microservices_template_golang/payment/src/eventmanager"
	"os"
	"os/signal"
	"syscall"
)

var (
	brokers             = []string{"kafka:9090"}
	topic   goka.Stream = "example-stream"
	group   goka.Group  = "example-group"
)

// process messages until ctrl-c is pressed
func runProcessor() {
	// process callback is invoked for each message delivered from
	// "example-stream" topic.
	cb := func(ctx goka.Context, msg interface{}) {
		log.Println("******** running processor ********")
		// SetValue stores the incremented counter in the group table for in
		// the message's key.
		//ctx.SetValue(counter)
		log.Printf("msg = %v",msg)
	}


	pc := eventmanager.NewConfig()
	cc := eventmanager.NewConfig()
	// Define a new processor group. The group defines all inputs, outputs, and
	// serialization formats. The group-table topic is "example-group-table".
	p, err := goka.NewProcessor(brokers,
		goka.DefineGroup(group,
		goka.Input(topic, new(codec.String), cb),
	),
		goka.WithProducerBuilder(kafka.ProducerBuilderWithConfig(pc)), // our config, mostly default
		goka.WithConsumerBuilder(kafka.ConsumerBuilderWithConfig(cc)), // our config, mostly default
	)
	if err != nil {
		log.Fatalf("error creating processor: %v", err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan bool)
	go func() {
		defer close(done)
		fmt.Println("******** run processor xzd")
		if err = p.Run(ctx); err != nil {
			log.Fatalf("error running processor: %v", err)
		}
	}()

	wait := make(chan os.Signal, 1)
	signal.Notify(wait, syscall.SIGINT, syscall.SIGTERM)
	<-wait   // wait for SIGINT/SIGTERM
	cancel() // gracefully stop processor
	<-done
}

func main() {
	//runEmitter()   // emits one message and stops
	runProcessor() // press ctrl-c to stop
}

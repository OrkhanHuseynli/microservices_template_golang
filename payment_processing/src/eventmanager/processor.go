package eventmanager

import (
	"context"
	"fmt"
	cluster "github.com/bsm/sarama-cluster"
	"github.com/lovoo/goka"
	"github.com/lovoo/goka/kafka"
	"log"
	"microservices_template_golang/payment/src/models"
	"os"
	"os/signal"
	"syscall"
)

type EventProcessor struct {
	processor *goka.Processor
}

func (e *EventProcessor) InitSimpleProcessor(brokers []string, group goka.Group, groupCallback func(ctx goka.Context, msg interface{}),
	topic goka.Stream, pc, cc *cluster.Config) {
	p, err := goka.NewProcessor(brokers,
		goka.DefineGroup(group,
			goka.Input(topic, new(models.PaymentCodec), groupCallback),
		),
		goka.WithProducerBuilder(kafka.ProducerBuilderWithConfig(pc)), // our config, mostly default
		goka.WithConsumerBuilder(kafka.ConsumerBuilderWithConfig(cc)), // our config, mostly default
	)

	if err != nil {
		log.Fatalf("error creating processor: %v", err)
	}

	e.processor = p
}

func (e *EventProcessor) InitDefaultProcessor(brokers []string, group goka.Group, topic goka.Stream){
	pc := NewConfig()
	cc := NewConfig()

	cb := func(ctx goka.Context, msg interface{}) {
		log.Printf("msg = %v", msg)
		payment, ok := msg.(*models.Payment)
		if !ok {
			log.Println("Error while parsing message to the structure")
		}
		log.Printf("Payment from %v was just processed", payment.Author)
	}

	e.InitSimpleProcessor(brokers, group, cb, topic, pc, cc)
}

func (e *EventProcessor) Run(){
	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan bool)
	go func() {
		defer close(done)
		fmt.Println("******** run processor xzd")
		if err := e.processor.Run(ctx); err != nil {
			log.Fatalf("error running processor: %v", err)
		}
	}()

	wait := make(chan os.Signal, 1)
	signal.Notify(wait, syscall.SIGINT, syscall.SIGTERM)
	<-wait   // wait for SIGINT/SIGTERM
	cancel() // gracefully stop processor
	<-done
}

package main


import (
	"context"
	"fmt"
	"github.com/Shopify/sarama"
	cluster "github.com/bsm/sarama-cluster"
	"github.com/lovoo/goka/kafka"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/lovoo/goka"
	"github.com/lovoo/goka/codec"
)

var (
	brokers             = []string{"kafka:9090"}
	topic   goka.Stream = "example-stream"
	group   goka.Group  = "example-group"
)


func NewConfig() *cluster.Config {
	c := kafka.NewConfig()
	c.Version = sarama.V2_1_0_0
	c.Consumer.Offsets.CommitInterval = 100 * time.Millisecond
	c.Producer.Return.Errors = true
	c.Producer.Return.Successes = true
	c.Producer.RequiredAcks = sarama.WaitForAll
	c.Producer.Compression = sarama.CompressionSnappy
	c.Producer.Flush.Messages = 16
	c.Producer.Flush.Frequency = 5 * time.Millisecond
	return c
}


// emits a single message and leave
func runEmitter() {
	pc := NewConfig()
	fmt.Println("\n*********************\n", brokers[0], "\n*********************\n")
	emitter, err := goka.NewEmitter(brokers, topic, new(codec.String), goka.WithEmitterProducerBuilder(kafka.ProducerBuilderWithConfig(pc)))
	fmt.Println(emitter)
	if err != nil {
		log.Fatalf("error creating emitter: %v", err)
	}
	defer emitter.Finish()
	err = emitter.EmitSync("some-key", "some-value")
	if err != nil {
		log.Fatalf("error emitting message: %v", err)
	}
	fmt.Println("message emitted")
}

// process messages until ctrl-c is pressed
func runProcessor() {
	// process callback is invoked for each message delivered from
	// "example-stream" topic.
	cb := func(ctx goka.Context, msg interface{}) {
		fmt.Println("******** callback")
		var counter int64
		// ctx.Value() gets from the group table the value that is stored for
		// the message's key.
		if val := ctx.Value(); val != nil {
			fmt.Println("*** Value ", val)
			counter = val.(int64)
		}
		counter++
		// SetValue stores the incremented counter in the group table for in
		// the message's key.
		ctx.SetValue(counter)
		log.Printf("key = %s, counter = %v, msg = %v", ctx.Key(), counter, msg)
	}


	pc := NewConfig()
	cc := NewConfig()
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
	runEmitter()   // emits one message and stops
	runProcessor() // press ctrl-c to stop
}

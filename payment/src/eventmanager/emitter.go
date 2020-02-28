package eventmanager

import (
	"github.com/Shopify/sarama"
	cluster "github.com/bsm/sarama-cluster"
	"github.com/lovoo/goka"
	"github.com/lovoo/goka/kafka"
	"log"
	"time"
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

func NewAppEmitter(brokers []string, topic goka.Stream, codec goka.Codec, producerConfig *cluster.Config) *goka.Emitter {
	emitter, err := goka.NewEmitter(brokers, topic, codec, goka.WithEmitterProducerBuilder(kafka.ProducerBuilderWithConfig(producerConfig)))
	if err != nil {
		log.Fatalf("error creating emitter: %v", err)
	}
	return emitter
}

func DefaultEmitter(brokers []string, topic goka.Stream, codec goka.Codec) *goka.Emitter {
	return NewAppEmitter(brokers, topic, codec, NewConfig())
}

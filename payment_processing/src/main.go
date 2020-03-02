package main

import (
	"github.com/lovoo/goka"
	"microservices_template_golang/payment_processing/src/eventmanager"
)

var (
	brokers             = []string{"kafka:9090"}
	topic   goka.Stream = "example-stream"
	group   goka.Group  = "example-group"
)

func main() {
	p := eventmanager.EventProcessor{}
	p.InitDefaultProcessor(brokers, group, topic)
	p.Run()
}

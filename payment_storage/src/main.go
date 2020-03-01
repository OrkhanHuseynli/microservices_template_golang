package main

import (
	"github.com/lovoo/goka"
	"microservices_template_golang/payment/src/eventmanager"
)

var (
	brokers             = []string{"kafka:9090"}
	topic   goka.Stream = "payment-storage"
	group   goka.Group  = "example-group"
)

func main() {
	p := eventmanager.EventProcessor{}
	p.InitDefaultProcessor(brokers, group, topic)
	p.Run()
}

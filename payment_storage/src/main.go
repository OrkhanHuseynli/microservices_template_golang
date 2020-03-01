package main

import (
	"microservices_template_golang/payment/src/core"
)

func main() {
	app := core.New()
	app.Start()
}

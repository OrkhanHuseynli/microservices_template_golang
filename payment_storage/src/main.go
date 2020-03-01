package main

import (
	"microservices_template_golang/payment_storage/src/core"
)

func main() {
	app := core.New()
	app.Start()
}

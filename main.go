package main

import (
	"learn-nats/publisher"
	"learn-nats/replier"
	"time"
)

func main() {
	go replier.Reply()
	time.Sleep(time.Second)

	go publisher.Publisher()

	select {}
}

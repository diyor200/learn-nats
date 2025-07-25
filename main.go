package main

import (
	"learn-nats/publisher"
	"learn-nats/subscriber"
	"time"
)

func main() {
	ready := make(chan struct{})
	ready2 := make(chan struct{})

	go subscriber.Subscriber(ready)
	go subscriber.Subscriber2(ready2)
	time.Sleep(time.Second)
	go publisher.Publisher()

	select {}
}

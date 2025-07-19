package main

import (
	"learn-nats/publisher"
	"learn-nats/subscriber"
	"time"
)

func main() {
	ready := make(chan struct{})

	go subscriber.Subscribe(ready)
	time.Sleep(time.Second)
	go publisher.Publisher()

	select {}
}

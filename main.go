package main

import (
	"learn-nats/publisher"
	"learn-nats/subscriber"
	"time"
)

func main() {
	go subscriber.Subscribe()
	time.Sleep(500 * time.Millisecond)
	go publisher.Publisher()

	select {}
}

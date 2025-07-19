package main

import (
	"learn-nats/publisher"
	"learn-nats/subscriber"
	"time"
)

func main() {
	go subscriber.Subscribe()
	time.Sleep(time.Second)
	go publisher.Publisher()

	select {}
}

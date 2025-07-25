package subscriber

import (
	"fmt"
	"time"

	"github.com/nats-io/nats.go"
)

func Subscriber(ready chan<- struct{}) {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		panic(err)
	}
	defer nc.Close()

	// declare jetstream
	js, err := nc.JetStream()
	if err != nil {
		panic(err)
	}

	// add stream
	streaminfo, err := js.AddStream(&nats.StreamConfig{
		Name:     "EVENTS",
		Subjects: []string{"orders.created", "orders.cancelled"},
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(streaminfo.Config.Discard)

	// declare dubscribers
	_, err = js.QueueSubscribe("orders.*", "orders", func(msg *nats.Msg) {
		fmt.Printf("worker1 received: %s\n", string(msg.Data))
		time.Sleep(time.Second)
		msg.Ack()
	}, nats.Durable("shared-durable"), nats.ManualAck())
	if err != nil {
		panic(err)
	}

	close(ready)

	fmt.Println("Listening [updates] subject ...")
	select {}
}

func Subscriber2(ready chan struct{}) {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		panic(err)
	}
	defer nc.Close()

	// declare jetstream
	js, err := nc.JetStream()
	if err != nil {
		panic(err)
	}

	// add stream
	streaminfo, err := js.AddStream(&nats.StreamConfig{
		Name:     "EVENTS",
		Subjects: []string{"orders.created", "orders.cancelled"},
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(streaminfo.Config.Discard)

	// declare dubscribers
	_, err = js.QueueSubscribe("orders.*", "orders", func(msg *nats.Msg) {
		fmt.Printf("worker2 received: %s\n", string(msg.Data))
		msg.Ack()
	}, nats.Durable("shared-durable"), nats.ManualAck())
	if err != nil {
		panic(err)
	}

	close(ready)

	fmt.Println("Listening [updates] subject ...")
	select {}
}

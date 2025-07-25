package publisher

import (
	"fmt"
	"time"

	"github.com/nats-io/nats.go"
)

func Publisher() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		panic(err)
	}
	defer nc.Close()

	fmt.Println("publisher connected ...")

	js, err := nc.JetStream()
	if err != nil {
		panic(err)
	}

	_, err = js.AddStream(&nats.StreamConfig{
		Name:     "EVENTS",
		Subjects: []string{"oreders.*"},
	})
	if err != nil && err != nats.ErrStreamNameAlreadyInUse {
		panic(err)
	}

	for i := 0; ; i++ {
		time.Sleep(300 * time.Millisecond)

		msg := fmt.Sprintf("Order ID: %d", i)
		_, err = js.Publish("orders.created", []byte(msg))
		if err != nil {
			panic(err)
		}

		msg = fmt.Sprintf("Cancelled order ID: %d", i)
		_, err = js.Publish("orders.cancelled", []byte(msg))
		if err != nil {
			panic(err)
		}
	}
}

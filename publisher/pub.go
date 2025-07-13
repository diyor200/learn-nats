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

	for {
		time.Sleep(time.Millisecond * 300)

		err = nc.Publish("updates", []byte("hello world"))
		if err != nil {
			panic(err)
		}

		fmt.Println("message published!")
	}
}

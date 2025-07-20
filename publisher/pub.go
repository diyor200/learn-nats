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

		msg, err := nc.Request("req-reply", []byte("hello world"), time.Second)
		if err != nil {
			panic(err)
		}

		fmt.Println("message published! Got reply: ", string(msg.Data))
	}
}

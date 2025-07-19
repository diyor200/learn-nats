package subscriber

import (
	"fmt"

	"github.com/nats-io/nats.go"
)

func Subscribe() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		panic(err)
	}

	fmt.Println("subscriber connected ...")

	defer nc.Close()

	_, err = nc.Subscribe("updates", func(msg *nats.Msg) {
		fmt.Printf("Received message from nats: %s\n", string(msg.Data))
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("Listening [updates] subject ...")
}

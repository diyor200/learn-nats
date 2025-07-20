package subscriber

import (
	"fmt"

	"github.com/nats-io/nats.go"
)

func Subscribe(ready chan<- struct{}) {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		panic(err)
	}

	fmt.Println("subscriber connected ...")

	defer nc.Close()

	// subscribers
	for i := range 5 {
		_, err = nc.QueueSubscribe("updates", "test.queue", func(msg *nats.Msg) {
			fmt.Printf("worker%d received: %s\n", i, string(msg.Data))
		})
		if err != nil {
			panic(err)
		}
	}

	close(ready)

	fmt.Println("Listening [updates] subject ...")
	select {}
}

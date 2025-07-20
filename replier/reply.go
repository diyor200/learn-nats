package replier

import (
	"fmt"

	"github.com/nats-io/nats.go"
)

func Reply() {
	ns, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		panic(err)
	}

	defer ns.Close()

	_, err = ns.Subscribe("req-reply", func(msg *nats.Msg) {
		response := "hello" + string(msg.Data)
		msg.Respond([]byte(response))
	})
	if err != nil {
		panic(err)
	}

	fmt.Println("Listening for requests on [greet] subject...")
	select {}
}

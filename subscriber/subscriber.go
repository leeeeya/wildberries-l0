package main

import (
	"fmt"
	stan "github.com/nats-io/stan.go"
	"log"
	"net/http"
)

func main() {
	sc, err := stan.Connect("test-cluster", "consumer")
	if err != nil {
		fmt.Println(err)
	}

	_, err = sc.Subscribe("foo", func(m *stan.Msg) {
		log.Printf("Received a message: %s\n", string(m.Data))
	})
	if err != nil {
		log.Println(err)
	}

	err = http.ListenAndServe("127.0.0.1:8081", nil)
	if err != nil {
		log.Println(err)
	}
}

package main

import (
	"fmt"
	stan "github.com/nats-io/stan.go"
	"log"
	"net/http"
)

type Stan struct {
	sc  stan.Conn
	msg *stan.Msg
}

func (s *Stan) postInfo(w http.ResponseWriter, r *http.Request) {
	w.Write(s.msg.Data)
}

func main() {
	sc, err := stan.Connect("test-cluster", "consumer")
	if err != nil {
		fmt.Println(err)
	}

	s := new(Stan)
	s.sc = sc

	_, err = sc.Subscribe("foo", func(m *stan.Msg) {
		s.msg = m
		http.HandleFunc("/wb_info", s.postInfo)
		log.Printf("Received a message: %s\n", string(m.Data))
	})
	if err != nil {
		log.Println(err)
	}

	err = http.ListenAndServe("127.0.0.1:8088", nil)
	if err != nil {
		log.Println(err)
	}
}

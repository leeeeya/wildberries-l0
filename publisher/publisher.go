package main

import (
	"encoding/json"
	stan "github.com/nats-io/stan.go"
	"log"
	"time"
	"wildberries/data"
	"wildberries/db"
)

func main() {

	db.InitDB()
	defer db.Close()

	data.GetData()
	db.FetchData()
	db.FetchCache()

	sc, err := stan.Connect("test-cluster", "publisher")
	if err != nil {
		log.Println(err)
	}
	if bytes, err := json.Marshal(data.Cache); err == nil {
		_ = sc.Publish("foo", bytes)
		time.Sleep(100 * time.Millisecond)
	}

}

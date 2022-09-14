package db

import (
	"context"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"time"
)

var DB *sql.DB

func ping(ctx context.Context) {
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	err := DB.PingContext(ctx)
	//println(err)
	if err != nil {
		log.Fatalf("unable to connect to database: %v", err)
	}
}

func InitDB() {
	connStr := "user=liya password=pg12345 dbname=wb sslmode=disable host=localhost port=5333"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Panic(err)
	}

	ctx, stop := context.WithCancel(context.Background())
	defer stop()
	DB = db
	ping(ctx)
}

//func InitDB() {
//	connection()
//}

func Close() {
	DB.Close()
	DB = nil
}

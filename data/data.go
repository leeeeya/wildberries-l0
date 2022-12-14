package data

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"time"
)

type Data struct {
	OrderUid    string `json:"order_uid,omitempty"`
	TrackNumber string `json:"track_number"`
	Entry       string `json:"entry"`
	Delivery    struct {
		Name    string `json:"name"`
		Phone   string `json:"phone"`
		Zip     string `json:"zip"`
		City    string `json:"city"`
		Address string `json:"address"`
		Region  string `json:"region"`
		Email   string `json:"email"`
	} `json:"delivery"`
	Payment struct {
		Transaction  string `json:"transaction"`
		RequestId    string `json:"request_id"`
		Currency     string `json:"currency"`
		Provider     string `json:"provider"`
		Amount       int    `json:"amount"`
		PaymentDt    int    `json:"payment_dt"`
		Bank         string `json:"bank"`
		DeliveryCost int    `json:"delivery_cost"`
		GoodsTotal   int    `json:"goods_total"`
		CustomFee    int    `json:"custom_fee"`
	} `json:"payment"`
	Items []struct {
		ChrtId      int    `json:"chrt_id"`
		TrackNumber string `json:"track_number"`
		Price       int    `json:"price"`
		Rid         string `json:"rid"`
		Name        string `json:"name"`
		Sale        int    `json:"sale"`
		Size        string `json:"size"`
		TotalPrice  int    `json:"total_price"`
		NmId        int    `json:"nm_id"`
		Brand       string `json:"brand"`
		Status      int    `json:"status"`
	} `json:"items"`
	Locale            string    `json:"locale"`
	InternalSignature string    `json:"internal_signature"`
	CustomerId        string    `json:"customer_id"`
	DeliveryService   string    `json:"delivery_service"`
	Shardkey          string    `json:"shardkey"`
	SmId              int       `json:"sm_id"`
	DateCreated       time.Time `json:"date_created"`
	OofShard          string    `json:"oof_shard"`
}

type MemCache map[string]Data

var D []Data

var Cache = make(MemCache)

var DataByte = make([][]byte, 0)

func openFile(fileName string) {
	var tmpD Data

	file, err := os.Open(fileName)
	if err != nil {
		log.Panic()
	}
	defer file.Close()

	tmp, err := io.ReadAll(file)
	if err != nil {
		log.Panic(err)
	} else if len(tmp) == 0 {
		log.Printf("%s: Empty file", fileName)
		return
	}

	if err = json.Unmarshal(tmp, &tmpD); err != nil {
		log.Panic(err)
	} else if tmpD.OrderUid == "" {
		log.Printf("%s: Invalid data", fileName)
		return
	}

	DataByte = append(DataByte, tmp)
	D = append(D, tmpD)
}

func GetData() {
	openFile("./models/model.json")
	openFile("./models/model1.json")
	openFile("./models/model2.json")
	openFile("./models/model3.json")
	openFile("./models/model4.json")
}

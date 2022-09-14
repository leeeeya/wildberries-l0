package db

import (
	"database/sql"
	"encoding/json"
	"log"
	"wildberries/data"
)

func FetchData() {
	_, err := DB.Exec(`INSERT INTO data(id, data) values ($1, $2)`, data.D.OrderUid, data.DataByte)
	if err != nil {
		log.Println(err)
		panic(err)
	}
}

func FetchCache() {
	var cacheData data.Data
	var tmp []byte
	var id string

	if rows, err := DB.Query(`SELECT * FROM wb.public.data`); err == nil {
		defer rows.Close()
		for rows.Next() == true {
			if err := rows.Scan(&id, &tmp); err != nil {
				log.Panic()
			} else {
				_ = json.Unmarshal(tmp, &cacheData)
				data.Cache[id] = cacheData
			}
		}
	} else if err != sql.ErrNoRows {
		log.Panic()
	}

}

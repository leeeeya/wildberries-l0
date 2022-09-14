package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"wildberries/data"
)

func FetchData() {
	allValues := make([]string, 0)

	for i := 0; i < len(data.D); i++ {
		value := fmt.Sprintf("('%s', '%s')", data.D[i].OrderUid, data.DataByte[i])
		allValues = append(allValues, value)
	}

	query := `INSERT INTO data(id, data) VALUES ` + strings.Join(allValues, ", ") +
		`on conflict do nothing`
	_, err := DB.Exec(query)
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

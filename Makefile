BIN_PATH=~/go/bin

service_up:
	docker-compose up -d

service_down:
	docker-compose down

db_up:
	$(BIN_PATH)/goose -dir db_migration postgres "user=liya password=pg12345 dbname=wb sslmode=disable host=localhost port=5333" up

db_down:
	$(BIN_PATH)/goose -dir db_migration postgres "user=liya password=pg12345 dbname=wb sslmode=disable host=localhost port=5333" down
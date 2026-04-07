include .env
export

service-run:
	go run main.go

migrate-up:
	migrate -path migration -database ${CONN_STR} up

migrate-down:
	migrate -path migration -database ${CONN_STR} down

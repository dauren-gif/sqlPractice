include .env
export

service-run:
	@go run main.go

service-deploy:
	docker compose up -d application

service-undeploy:
	docker compose down application

migrate-up:
	migrate -path migration -database ${CONN_STR} up

migrate-down:
	migrate -path migration -database ${CONN_STR} down

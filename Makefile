broker-up:
	docker compose up -d

broker-down:
	docker compose down

processor:
	go run cmd/processor/main.go -balance -above_threshold

api:
	go run cmd/service/main.go

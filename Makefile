broker:
	docker compose up -d

processor:
	go run cmd/processor/main.go -balance -above_threshold

api:
	go run cmd/service/main.go

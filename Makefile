broker:
	docker compose up -d

processor:
	go run cmd/processor/main.go -balance -above_threshold

service:
	go run cmd/service/main.go

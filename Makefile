.PHONY: run
run:
	./bin/bot

.PHONY: build
build:
	go build -o bin/bot cmd/bot/main.go

.PHONY: test
test:
	go test -race ./internal/...

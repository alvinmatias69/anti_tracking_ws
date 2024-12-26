.PHONY: all
all:
	go build -o ws -ldflags "-s -w" cmd/main.go

.PHONY: dev
dev:
	go run cmd/main.go

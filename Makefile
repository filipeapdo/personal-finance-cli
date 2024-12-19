.DEFAULT_GOAL := run

clean:
	go clean ./...

fmt:
	go fmt ./...

vet: fmt
	go vet ./...

run: clean vet
	go run main.go

test:
	go test ./...

.PHONY:fmt vet run test

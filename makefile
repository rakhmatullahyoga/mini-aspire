.PHONY: dep tidy compile test run

run: compile
	./mini-aspire

dep:
	go mod download

tidy:
	go mod tidy

compile:
	go build -o mini-aspire cmd/main.go

test:
	go test ./... -cover

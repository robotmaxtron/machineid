.PHONY: tidy build clean default test

build: clean
	@go build -o machineid ./cmd/machineid/main.go

clean:
	@rm -f ./machineid

tidy:
	go mod tidy

test:
	go test

default: tidy build

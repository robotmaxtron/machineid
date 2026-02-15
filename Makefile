.PHONY: build clean default test

build: clean
	@go build -o machineid ./cmd/machineid/main.go

clean:
	@rm -f ./machineid

test:
	go test

default: build

.PHONY: build run clean

build:
	@ go build -o bin/main cmd/api/main.go

run: build
	@ ./bin/main

clean:
	@ rm -rf bin

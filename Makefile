.PHONY: build clean

build:
	env GOOS=linux go build -ldflags="-s -w" -o bin/WebErrors WebErrors/main.go

clean:
	rm -rf ./bin

BINARY_NAME=hello-world

build:
	go build -o bin/${BINARY_NAME} main.go

run:
 bin/${BINARY_NAME}

build-and-run: build run

cross-compile-amd64:
	GOOS=darwin GOARCH=amd64 go build -o bin/${BINARY_NAME}-macos-amd64 main.go
	GOOS=linux GOARCH=amd64 go build -o bin/${BINARY_NAME}-linux-amd64 main.go
	GOOS=windows GOARCH=amd64 go build -o bin/${BINARY_NAME}-windows-amd64 main.go

cross-compile-386:
	GOOS=darwin  GOARCH=386 go build -o bin/${BINARY_NAME}-macos-386 main.go
	GOOS=linux   GOARCH=386 go build -o bin/${BINARY_NAME}-linux-386 main.go
	GOOS=windows GOARCH=386 go build -o bin/${BINARY_NAME}-windows-386 main.go

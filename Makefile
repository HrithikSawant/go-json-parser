build:
	go build -o bin/go-json-parser main.go

test:
	go test ./...

run:
	go run main.go myfile.json

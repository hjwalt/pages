tidy:
	go mod tidy
	go fmt ./...

update:
	go get -u ./...
	go mod tidy

build:
	go build -o bin/ui

proto:
	clang-format -style=file:.clang-format -i **/*.proto
	protoc  --proto_path=. --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative --go_out=. --go-grpc_out=. ./**/*.proto

run: 
	go run ./main
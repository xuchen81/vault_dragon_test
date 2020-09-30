compile:
	GOOS=linux GOARCH=amd64 go build -o bin/vaultdragon-api-linux-amd64 main.go

build:
	GOOS=linux GOARCH=amd64 go build -o bin/vaultdragon-api-linux-amd64 main.go
	GOOS=darwin GOARCH=amd64 go build -o bin/vaultdragon-api-darwin-amd64 main.go
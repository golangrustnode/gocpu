all:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o cpuinfo cmd/main.go
build:
	go build -o $(GOPATH)/bin/tb-logparser cmd/tb-logparser/main.go
	go build -o $(GOPATH)/bin/tb-maillog cmd/tb-maillog/main.go
run-parser:
	go run cmd/tb-logparser/main.go
run-maillog:
	go run cmd/tb-maillog/main.go
compile:
	# MacOS
	GOOS=darwin GOARCH=amd64 go build -o bin/tb-logparser-mac-amd64 cmd/tb-logparser/main.go
	# Linux
	GOOS=linux GOARCH=amd64 go build -o bin/tb-logparser-linux-amd64 cmd/tb-logparser/main.go
run: 
	go run main.go

test: 
	go test ./... -v -gcflags=-l

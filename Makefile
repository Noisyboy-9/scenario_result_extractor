status: 
	go run main.go status

build: 
	go build -o ./build/project main.go 

test: 
	go test ./... -v -gcflags=-l

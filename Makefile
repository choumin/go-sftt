GOPATH:=$(CURDIR)
export GOPATH

all: client server

client:
	go build -o cmd/client internal/app/client/main.go
server:
	go build -o cmd/server internal/app/server/main.go

clean:
	rm -rf cmd/client
	rm -rf cmd/server

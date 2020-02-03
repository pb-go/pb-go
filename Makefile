.PHONY: build generate test clean server client
build:generate server client
generate:
	go generate main/main.go
test:
	go test -coverprofile=output/coverage.out ./...
clean:
	rm -rf output
	rm -rf templates/*.qtpl.go
server:
	go build -o output/pb-go main/main.go
client:
	# todo: build pg-go-cli

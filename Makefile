.PHONY: build test clean
build:
	go build -o output/pb-go main/main.go
test:
	go test -coverprofile=output/coverage.out ./...
clean:
	rm -rf output

.PHONY: build test clean
build:generate
	go build -o output/pb-go main/main.go
generate:
	go generate main/main.go
test:
	go test -coverprofile=output/coverage.out ./...
clean:
	rm -rf output
	rm -rf templates/*.qtpl.go

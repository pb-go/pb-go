.PHONY: build generate test clean server client
build:generate server client
generate:
	go generate cmd/server/main.go
test:
	go test -coverprofile=output/coverage.out ./...
clean:
	rm -rf output
	rm -rf templates/*.qtpl.go
server:
	go build -o output/pb-go cmd/server/main.go
client:
	go build -o output/pb-cli cmd/client/main.go
all-platform:generate
	OUTPUT=./output/pb-go ./scripts/build-all.sh ./cmd/server/main.go
	OUTPUT=./output/pb-cli ./scripts/build-all.sh ./cmd/client/main.go
	bash ./scripts/archive-release.sh
	bash ./scripts/sha256sums.sh
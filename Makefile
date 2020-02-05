.PHONY: build generate test clean server client
build: showcpu clean generate server client
showcpu:
	cat /proc/cpuinfo
	cat /proc/meminfo
generate:
	go generate cmd/server/main.go
	ls -alh ./statik/
test:
	go test -coverprofile=output/coverage.out ./...
clean:
	rm -rf output
	rm -rf templates/*.qtpl.go
	rm -rf statik
	rm -rf static/*.fasthttp.gz
server:
	go build -race -o output/pb-go cmd/server/main.go
	ls -alh ./output/
client:
	go build -race -o output/pb-cli cmd/client/main.go
	ls -alh ./output/
all-platform:generate
	OUTPUT=./output/pb-go ./scripts/build-all.sh ./cmd/server/main.go
	OUTPUT=./output/pb-cli ./scripts/build-all.sh ./cmd/client/main.go
	strip ./output/pb-go
	strip ./output/pb-cli
	bash ./scripts/archive-release.sh
	bash ./scripts/sha256sums.sh

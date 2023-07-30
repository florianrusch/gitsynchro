BINARY_NAME=gitsynchro

build:
	go build -o ${BINARY_NAME} main.go

run: build
	./${BINARY_NAME}

clean:
	go clean
	rm ${BINARY_NAME} | true

test:
	go test ./...

test_coverage:
	go test ./... -coverprofile=coverage.out

dep:
	go mod download

vet:
	go vet

lint:
	go mod tidy
	go run github.com/golangci/golangci-lint/cmd/golangci-lint@v1.53.3 run --enable-all

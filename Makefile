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
	golangci-lint run --enable-all

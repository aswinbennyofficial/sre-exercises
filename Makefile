BINARY_NAME=app-name
BUILD_DIR=output/bin

build:
	GOARCH=amd64 GOOS=linux go build -o ./${BUILD_DIR}/${BINARY_NAME}-linux ./cmd/main/main.go
	GOARCH=amd64 GOOS=darwin go build -o ./${BUILD_DIR}/${BINARY_NAME}-darwin ./cmd/main/main.go
	GOARCH=amd64 GOOS=windows go build -o ./${BUILD_DIR}/${BINARY_NAME}-windows ./cmd/main/main.go
	@echo "Executables are created in ${BUILD_DIR}"


run: dep 
	go run ./cmd/main/

clean:
	go clean
	rm -rf ${BUILD_DIR}

test:
	go test ./...

test_coverage:
	go test ./... -coverprofile=coverage.out

dep:
	go mod download

vet:
	go vet

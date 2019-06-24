GO_BIN := ./bin
EXE_PATH := $(GO_BIN)/perf-test

fmt:
	go fmt ./...

dep:
	go mod tidy && go mod vendor && go mod verify

build: mkdirs
	go build -o $(EXE_PATH) ./cmd/main.go

mkdirs:
	mkdir -p $(GO_BIN)

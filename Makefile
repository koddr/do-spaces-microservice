.PHONY: clean test security build run

APP_NAME = do_spaces
BUILD_DIR = $(PWD)/build

clean:
	rm -rf ./build

security:
	gosec -quiet ./...

test: security
	go test -cover ./...

build: clean test
	CGO_ENABLED=0 go build -ldflags="-w -s" -o $(BUILD_DIR)/$(APP_NAME) main.go

run: build
	$(BUILD_DIR)/$(APP_NAME)
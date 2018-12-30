BUILD_ENVPARMS:=CGO_ENABLED=0
BIN?=./bin/url2data
VGO_EXEC:=go
export GO111MODULE=on


# install project dependencies
.PHONY: .deps
deps:
	$(info #Install dependencies...)
	$(VGO_EXEC) mod download

# run unit tests
.PHONY: .test
test: deps
	$(info #Running tests...)
	$(VGO_EXEC) test ./...

.PHONY: .fast-build
fast-build: deps
	$(info #Building...)
	$(BUILD_ENVPARMS) $(VGO_EXEC) build -ldflags "$(LDFLAGS)" -o $(BIN) ./*.go

.PHONY: .build
build: test fast-build

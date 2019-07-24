TEST_DIR := $(PWD)/tests
TOOLS_DIR := $(PWD)/tools

COVERAGE_PROFILE_FILE := $(TEST_DIR)/c.out
COVERAGE_REPORT_FILE := $(TEST_DIR)/c.html

ORIGIN := github.com
APP_NAME := studtool

LIB_NAME := go-rest
LIB_PACKAGE := $(ORIGIN)/$(APP_NAME)/$(LIB_NAME)

TEST_SERVER_ADDRESS ?= 127.0.0.1:80
LD_FLAGS := -X $(LIB_PACKAGE)/pkg/rest.testServerAddress=$(TEST_SERVER_ADDRESS)

TEST_FLAGS := -mod=vendor -v -race -coverprofile='$(COVERAGE_PROFILE_FILE)' -covermode=atomic -ldflags='$(LD_FLAGS)'
COVER_FLAGS := -html='$(COVERAGE_PROFILE_FILE)' -o '$(COVERAGE_REPORT_FILE)'

fmt:
	go fmt ./...

dep:
	go mod tidy && go mod vendor && go mod verify

gen:
	PATH='$(PATH):$(TOOLS_DIR)' go generate ./...

test: mkdirs bootstrap dep
	go test $(TEST_FLAGS) ./...
	go tool cover $(COVER_FLAGS)
	./upload_coverage.sh '$(COVERAGE_PROFILE_FILE)'

clean:
	rm '$(COVERAGE_PROFILE_FILE)'
	rm '$(COVERAGE_REPORT_FILE)'

lint: mkdirs bootstrap
	'$(TOOLS_DIR)/golangci-lint' run

mkdirs:
	mkdir -p '$(TEST_DIR)'
	mkdir -p '$(TOOLS_DIR)'

bootstrap:
	./install_tools.sh '$(TOOLS_DIR)'

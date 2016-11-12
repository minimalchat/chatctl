# Commands
GO_CMD = go
LINT_CMD = $(GOPATH)/bin/golint

# Directories
PACKAGE = github.com/minimalchat/chatctl
SRC = $(GOPATH)/src/$(PACKAGE)
DIST = $(GOPATH)/bin

.PHONY: lint build

build: lint test coverage clean compile

run: lint test go

lint:
	$(LINT_CMD) $(SRC)
	# $(LINT_CMD) $(SRC) $(TEST)

test:
	cd $(SRC)
	$(GO_CMD) test -v ./...
	# $(GOPATH)/bin/goveralls -coverprofile=coverage.out -service=travis-ci -repotoken $(COVERALLS_TOKEN)

coverage:
	cd $(SRC)
	$(DIST)/overalls -project=$(PACKAGE) -covermode=count
	$(GOPATH)/bin/goveralls -coverprofile=overalls.coverprofile -service=travis-ci -repotoken $(COVERALLS_TOKEN)

clean:
	rm -rf $(DIST)/chatctl

compile:
	mkdir -p $(DIST)
	cd $(SRC)
	$(GO_CMD) build -o $(DIST)/chatctl

go:
	cd $(SRC)
	$(GO_CMD) run main.go

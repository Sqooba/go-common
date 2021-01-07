
GOOS=linux
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOLINT=golangci-lint run

ensure:
	GOOS=${GOOS} $(GOCMD) mod vendor

clean:
	$(GOCLEAN) ./...

lint:
	$(GOLINT) ./...

build:
	GOOS=${GOOS} $(GOBUILD) ./...

test:
	$(GOCMD) test ./...

release:
	$(GOCMD) install ./...
	GOOS=${GOOS} $(GOCMD) install ./...

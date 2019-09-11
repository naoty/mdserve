VERSION ?= $$(git describe --tags)

build:
	go build -ldflags "-X main.Version=$(VERSION)" -o bin/mdserve

install: build
	cp bin/mdserve $(GOPATH)/bin/mdserve

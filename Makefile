CWD=$(shell pwd)
GOPATH := $(CWD)

build:	rmdeps deps fmt bin

prep:
	if test -d pkg; then rm -rf pkg; fi

self:   prep
	if test -d src/github.com/whosonfirst/go-whosonfirst-index; then rm -rf src/github.com/whosonfirst/go-whosonfirst-index; fi
	mkdir -p src/github.com/whosonfirst/go-whosonfirst-index
	cp -r *.go src/github.com/whosonfirst/go-whosonfirst-index/
	if test -d vendor/src; then cp -r vendor/src/* src/; fi

rmdeps:
	if test -d src; then rm -rf src; fi 

deps:

vendor-deps: rmdeps deps
	if test ! -d vendor; then mkdir vendor; fi
	if test -d vendor/src; then rm -rf vendor/src; fi
	cp -r src vendor/src
	find vendor -name '.git' -print -type d -exec rm -rf {} +
	rm -rf src

fmt:
	go fmt cmd/*.go
	go fmt *.go

bin:	self
	@GOPATH=$(GOPATH) go build -o bin/wof-index-count cmd/wof-index-count.go	

all: run

deps:
	go get -d -v ./...

updatedeps:
	go get -d -v -u -f ./...

build: deps
	go build ./...

run: deps env
	go run main.go main.pb.go

clean:
	go clean ./...

proto:
	go get -v go.pedge.io/protoeasy/cmd/protoeasy
	go get -v go.pedge.io/pkg/cmd/strip-package-comments
	protoeasy --go .
	find . -name *\.pb\*\.go | xargs strip-package-comments

env:
ifndef CURRENT_TOKEN
	$(error CURRENT_TOKEN must be set)
endif
ifndef CURRENT_SYSLOG_ADDRESS
	$(error CURRENT_SYSLOG_ADDRESS must be set)
endif

.PHONY: \
	all \
	deps \
	updatedeps \
	build \
	run \
	clean \
	proto \
	env

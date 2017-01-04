.PHONY: build clean

BINARY=mesos-tail

build: ${BINARY}

.get-deps: *.go
	go get -t -d -v ./...
	touch .get-deps

test: *.go .get-deps
	go test -v -cover "./..."

clean:
	rm -f .get-deps
	rm -f ${BINARY}

${BINARY}: *.go .get-deps
	go build -o $@ *.go


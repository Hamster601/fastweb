BINARY="fastweb"
VERSION=0.0.1
BUILD=`date+%FT%T%z`

PACKAGES=`golist./...|grep-v/vendor/`
VETPACKAGES=`golist./...|grep-v/vendor/|grep-v/examples/`
GOFILES=`find.-name"*.go"-typef-not-path"./vendor/*"`

default:
@go build -o ${BINARY} -tags=jsoniter

list:
@echo ${PACKAGES}
@echo ${VETPACKAGES}
@echo ${GOFILES}

fmt:
@go fmt -s -w ${GOFILES}

install:
@govendorsync-v

test:
@go test -cpu=1,2,4 -v -tags integration./...

vet:
@go vet $(VETPACKAGES)

docker:
@docker build -t wuxiaoxiaoshen/example:latest.

clean:
@if[-f${BINARY}];thenrm${BINARY};fi

.PHONY:default fmt fmt-check install test vet docker clean




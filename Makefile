BINARY="fastweb"
VERSION=1.0.0
BUILD=`date+%FT%T%z`

PACKAGES=`golist./...|grep-v/vendor/`
VETPACKAGES=`golist./...|grep-v/vendor/|grep-v/examples/`
GOFILES=`find.-name"*.go"-typef-not-path"./vendor/*"`

default:
@gobuild-o${BINARY}-tags=jsoniter

list:
@echo${PACKAGES}
@echo${VETPACKAGES}
@echo${GOFILES}

fmt:
@gofmt-s-w${GOFILES}

fmt-check:
@diff=?(gofmt-s-d$(GOFILES));
if[-n"$$diff"];then
echo"Pleaserun'makefmt'andcommittheresult:";
echo"$${diff}";
exit1;
fi;

install:
@govendorsync-v

test:
@gotest-cpu=1,2,4-v-tagsintegration./...

vet:
@govet$(VETPACKAGES)

docker:
@dockerbuild-twuxiaoxiaoshen/example:latest.

clean:
@if[-f${BINARY}];thenrm${BINARY};fi

.PHONY:defaultfmtfmt-checkinstalltestvetdockerclean




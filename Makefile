GOFMT=gofmt
GC=go build
VERSION = $(shell git describe --abbrev=4 --always --tags)
BUILD_NODE_PAR = -ldflags "-X github.com/NaturalChain/natural/common/config.Version=$(VERSION)"

SRC_FILES = $(shell git ls-files | grep -e .go$ | grep -v _test.go)

natural: $(SRC_FILES)
	$(GC) $(BUILD_NODE_PAR) -o natural main.go

format:
	$(GOFMT) -w main.go

clean:
	rm -rf natural


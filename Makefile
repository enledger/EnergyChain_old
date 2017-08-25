all: get_vendor_deps install

build:
	go build ./...

install:
	go install ./cmd/...

get_vendor_deps:
	go get github.com/Masterminds/glide
	glide install

clean:
	# maybe cleaning up cache and vendor is overkill, but sometimes
	# you don't get the most recent versions with lots of branches, changes, rebases...
	@rm -rf ~/.glide/cache/src/https-github.com-tendermint-*
	@rm -rf ./vendor
	@rm -f $GOPATH/bin/{energychain,energycli}


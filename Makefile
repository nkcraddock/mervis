VENDOR = $(CURDIR)/_vendor
GOPATH := $(VENDOR):$(GOPATH)
SERVER_FILES := $(shell find cmd/mervis -type f -name "*.go" ! -name "*_test.go")

default: vendor build

run:
	cd $(CURDIR)
	go run $(SERVER_FILES) -r client/build

test:
	cd $(CURDIR)
	go test -v ./...

# download deps to _vendor (and remove the git repos)
vendor:
	GOPATH=$(VENDOR)
	mkdir -p $(VENDOR)
	go get -d github.com/gorilla/mux
	go get -d github.com/onsi/gomega
	go get -d github.com/onsi/ginkgo/ginkgo
	find $(VENDOR) -type d -name '.git' | xargs rm -rf

clean:
	rm -rf build/

build: clean client
	mkdir -p build/
	CGO_ENABLED=0 go build -a -installsuffix cgo -o build/mervis --ldflags '-s' $(SERVER_FILES)

client-deps:
	if [ ! -d "client/node_modules" ]; then \
		cd client; \
		npm install; \
		bower install; \
	fi;

client: client-deps
	GOPATH=$(VENDOR)
	mkdir -p $(VENDOR)
	go get github.com/jteeuwen/go-bindata/...
	grunt --gruntfile client/Gruntfile.js package
	$(VENDOR)/bin/go-bindata -o "./res/client/resources.go" -pkg="client" -prefix="client/build/" client/build/...

docker: 
	docker build -t nkcraddock/mervis .

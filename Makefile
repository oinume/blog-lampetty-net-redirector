GCP_PROJECT = blog-lampetty-net-redirector
GO_TEST ?= go test -v -race -p=1
GO_TEST_PACKAGES = ./...

.PHONY: run
run:
	dev_appserver.py appengine/app/app.yaml

.PHONY: deploy
deploy:
	rm -f `pwd`/appengine/gopath/vendor/src
	mkdir -p `pwd`/appengine/gopath/vendor
	ln -s `pwd`/vendor `pwd`/appengine/gopath/vendor/src
	GOPATH="appengine/gopath/vendor" gcloud app deploy appengine/app/app.yaml --project=$(GCP_PROJECT)

.PHONY: browse
browse:
	gcloud app browse --project=$(GCP_PROJECT)


.PHONY: build
build:
	go build -o server ./cmd/server

.PHONY: test
test:
	$(GO_TEST) $(GO_TEST_PACKAGES)

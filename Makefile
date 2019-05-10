GCP_PROJECT_ID = blog-lampetty-net-redirector
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
	gcloud app browse --project=$(GCP_PROJECT_ID)


.PHONY: build
build:
	CGO_ENABLED=0 go build -o server ./cmd/server

.PHONY: test
test:
	$(GO_TEST) $(GO_TEST_PACKAGES)

.PHONY: docker/build
docker/build:
	docker build --pull -f Dockerfile \
	--tag gcr.io/$(GCP_PROJECT_ID)/server:$(IMAGE_TAG) .

.PHONY: gcloud/builds
gcloud/builds:
	gcloud builds submit --project $(GCP_PROJECT_ID) \
  	--tag gcr.io/$(GCP_PROJECT_ID)/blog-lampetty-net-redirector:$(IMAGE_TAG)

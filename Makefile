GCP_PROJECT = blog-lampetty-net-redirector
GOPATH = appengine/gopath/vendor

.PHONY: run
run:
	dev_appserver.py appengine/app/app.yaml

.PHONY: deploy
deploy:
	rm -f `pwd`/appengine/gopath/vendor/src
	mkdir -p `pwd`/appengine/gopath/vendor
	ln -s `pwd`/vendor `pwd`/appengine/gopath/vendor/src
	GOPATH=$(GOPATH) gcloud app deploy appengine/app/app.yaml --project=$(GCP_PROJECT)

.PHONY: browse
browse:
	gcloud app browse --project=$(GCP_PROJECT)

package main

import (
	"fmt"
	"net/http"
	"regexp"

	"strings"

	"google.golang.org/appengine"
)

const urlPrefix = "https://oinume.hatenablog.com"

var (
	archivesRe = regexp.MustCompile(`^/blog_ja/index.php/archives/([0-9]+)$`)
	categoryRe = regexp.MustCompile(`^/blog_ja/index.php/archives/category/(.+)$`)
	feedRe     = regexp.MustCompile(`^/blog_ja/index.php/feed`)
	techFeedRe = regexp.MustCompile(`^/tech/index.php/feed`)
)

func main() {
	http.Handle("/", newMux())
	appengine.Main()
}

func newMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handle)
	return mux
}

func handle(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	if strings.HasPrefix(path, "/robots.txt") {
		http.Error(w, "Not found", http.StatusNotFound)
	} else if m := archivesRe.FindStringSubmatch(path); len(m) > 1 {
		http.Redirect(w, r, fmt.Sprintf("%s/entry/wp/%s", urlPrefix, m[1]), http.StatusMovedPermanently)
	} else if m := categoryRe.FindStringSubmatch(path); len(m) > 1 {
		http.Redirect(w, r, fmt.Sprintf("%s/category/%s", urlPrefix, m[1]), http.StatusMovedPermanently)
	} else if m := feedRe.FindStringSubmatch(path); len(m) > 0 {
		http.Redirect(w, r, fmt.Sprintf("%s/rss", urlPrefix), http.StatusMovedPermanently)
	} else if m := techFeedRe.FindStringSubmatch(path); len(m) > 0 {
		http.Redirect(w, r, fmt.Sprintf("%s/rss", urlPrefix), http.StatusMovedPermanently)
	} else {
		http.Redirect(w, r, urlPrefix+"/", http.StatusMovedPermanently)
	}
}

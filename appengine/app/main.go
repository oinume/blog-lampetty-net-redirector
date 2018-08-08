package main

import (
	"fmt"
	"net/http"
	"regexp"

	"google.golang.org/appengine"
)

const urlPrefix = "https://oinume.hatenablog.com"

func main() {
	mux := newMux()
	http.Handle("/", mux)
	appengine.Main()
}

func newMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/blog_ja/index.php/archives/category", category)
	mux.HandleFunc("/blog_ja/index.php/archives", archives)
	mux.HandleFunc("/blog_ja/index.php/feed", feed)
	mux.HandleFunc("/blog_ja", root)
	mux.HandleFunc("/tech/index.php/feed", feed)
	mux.HandleFunc("/tech", root)
	mux.HandleFunc("/", root)
	return mux
}

func root(w http.ResponseWriter, r *http.Request) {
	redirectToRoot(w, r)
}

func archives(w http.ResponseWriter, r *http.Request) {
	re := regexp.MustCompile(`^/blog_ja/index.php/archives/([0-9]+)$`)
	matches := re.FindStringSubmatch(r.URL.Path)
	if len(matches) > 1 {
		http.Redirect(w, r, fmt.Sprintf("%s/entry/wp/%s", urlPrefix, matches[1]), http.StatusMovedPermanently)
		return
	}
	redirectToRoot(w, r)
}

func category(w http.ResponseWriter, r *http.Request) {
	re := regexp.MustCompile(`^/blog_ja/index.php/archives/category/(.+)$`)
	matches := re.FindStringSubmatch(r.URL.Path)
	if len(matches) > 1 {
		http.Redirect(w, r, fmt.Sprintf("%s/category/%s", urlPrefix, matches[1]), http.StatusMovedPermanently)
		return
	}
	redirectToRoot(w, r)
}

func feed(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, fmt.Sprintf("%s/rss", urlPrefix), http.StatusMovedPermanently)
}

func redirectToRoot(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, urlPrefix+"/", http.StatusMovedPermanently)
}

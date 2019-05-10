package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

const urlPrefix = "https://journal.lampetty.net"

var (
	archivesRe = regexp.MustCompile(`^/blog_ja/index.php/archives/([0-9]+)$`)
	categoryRe = regexp.MustCompile(`^/blog_ja/index.php/archives/category/(.+)$`)
	feedRe     = regexp.MustCompile(`^/blog_ja/index.php/feed`)
	techFeedRe = regexp.MustCompile(`^/tech/index.php/feed`)
)

func main() {
	if err := run(); err != nil {
		log.Printf("failed to run: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	return http.ListenAndServe(fmt.Sprintf(":%v", port), newMux())
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

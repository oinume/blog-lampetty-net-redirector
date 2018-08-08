package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var server *httptest.Server

func TestMain(m *testing.M) {
	mux := newMux()
	server = httptest.NewServer(mux)
	status := m.Run()
	server.Close()
	os.Exit(status)
}

func TestRoot(t *testing.T) {
	tests := []struct {
		path       string
		statusCode int
		location   string
	}{
		{
			path:       "/",
			statusCode: http.StatusMovedPermanently,
			location:   "https://oinume.hatenablog.com/",
		},
		{
			path:       "/tech/",
			statusCode: http.StatusMovedPermanently,
			location:   "https://oinume.hatenablog.com/",
		},
		{
			path:       "/tech/index.php/feed?q=dummy",
			statusCode: http.StatusMovedPermanently,
			location:   "https://oinume.hatenablog.com/rss",
		},
	}

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	for _, test := range tests {
		resp, err := client.Get(server.URL + test.path)
		if err != nil {
			t.Fatal(err)
		}

		t.Logf("path = %v", test.path)
		if got, want := resp.StatusCode, http.StatusMovedPermanently; got != want {
			t.Errorf("unexpected StatusCode: got=%v, want=%v\n", got, want)
		}
		if got, want := resp.Header.Get("Location"), test.location; got != want {
			t.Errorf("unexpected Location: got=%v, want=%v\n", got, want)
		}
		resp.Body.Close()
	}
}

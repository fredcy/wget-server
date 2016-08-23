package main

// This simple HTTP server serves static files as obtained by "wget -r -p -k"
// for some site. The files created by wget have any query string in the request
// included literally in their name and so we handle the requested URLs the same
// way, treating the query string as part of the name of the static file to be
// found and served.

import (
	"flag"
	"io"
	"log"
	"net/http"
	"path/filepath"
)

var address = flag.String("address", ":3000", "Listen and serve at this address")
var base = flag.String("base", ".", "base directory for files")

func handler(fs http.Dir) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if r.URL.RawQuery != "" {
			path = path + "?" + r.URL.RawQuery
		}

		log.Printf("handle %v", r.URL)

		// net/http is often able to guess the content-type and add it
		// automatically, but we need to give it some help for some cases.
		suffix := filepath.Ext(r.URL.Path)
		if suffix == ".css" {
			w.Header().Set("Content-Type", "text/css")
		}

		f, err := fs.Open(path)
		if err != nil {
			log.Printf("Cannot open %v: %v", path, err)
			http.NotFound(w, r)
		} else {
			if _, err := io.Copy(w, f); err != nil {
				log.Printf("Cannot copy %v: %v", path, err)
				http.Error(w, err.Error(), 500)
			}
			f.Close()
		}
	}
}

func main() {
	flag.Parse()
	fs := http.Dir(*base)

	http.HandleFunc("/", handler(fs))
	log.Printf("Listening on %v ...", *address)
	http.ListenAndServe(*address, nil)
}

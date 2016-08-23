package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func handler(fs http.Dir) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if r.URL.RawQuery != "" {
			path = path + "?" + r.URL.RawQuery
		}

		log.Printf("handle %v", r.URL)

		f, err := fs.Open(path)
		if err != nil {
			log.Printf("Cannot open %v: %v", path, err)
			fmt.Fprintf(w, "Cannot open %v: %v", path, err)
		} else {
			http.ServeContent(w, r, "todo", time.Now(), f)
			f.Close()
		}
	}
}

func main() {
	fs := http.Dir("/Users/fcy/Downloads/imsaalumni.org")

	http.HandleFunc("/", handler(fs))
	log.Println("Listening...")
	http.ListenAndServe(":3000", nil)
}

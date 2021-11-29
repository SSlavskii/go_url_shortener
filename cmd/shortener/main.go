package main

import (
	"io"
	"net/http"
	"net/url"
	"path"
	"strconv"
)

var url_to_int = make(map[string]int)
var int_to_url = make([]string, 0)

func UrlHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		u, _ := url.Parse(r.URL.RequestURI())
		index, err := strconv.Atoi(path.Base(u.Path))
		if err != nil {
			http.Error(w, err.Error(), 400)
		} else if index >= len(int_to_url) {
			http.Error(w, "NO such id", 400)
		} else {
			w.Header().Add("Content-Type", "application/json")
			w.Header().Set("Location", int_to_url[index])
			w.WriteHeader(http.StatusTemporaryRedirect)

		}
	} else if r.Method == http.MethodPost {
		defer r.Body.Close()
		raw_url, err := io.ReadAll(r.Body)
		short_int, ok := url_to_int[string(raw_url)]
		if !ok {
			short_int = len(int_to_url)
			url_to_int[string(raw_url)] = short_int
			int_to_url = append(int_to_url, string(raw_url))
		}
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		w.Header().Add("Content-Type", "application/json")
		w.Header().Add("Accept-Charset", "utf-8")
		w.WriteHeader(201)
		w.Write([]byte("http://localhost:8080/" + strconv.Itoa(short_int)))
	} else {
		http.Error(w, "Only GET and POST requests are allowed!", 400)
	}

}

func main() {

	http.HandleFunc("/", UrlHandler)
	http.ListenAndServe(":8080", nil)
}

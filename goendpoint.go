package main

import (
    "fmt"
    "log"
    "net/http"
	"net/url"
)

func handler(w http.ResponseWriter, r *http.Request) {
    //fmt.Fprintf(w, "Hi there, I love %s!\n", r.URL.Path[1:])
	u, _ := url.Parse(r.URL.String())
	queryParams := u.Query()
	fmt.Fprintf(w, "Hi there, I love %s!\n", r.URL.Path)
	fmt.Fprintf(w, "Url query params is: %s\n", queryParams)
	fmt.Fprintf(w, "Not formated query: %s\n", u)
	for k, v := range r.Header {
        fmt.Fprintf(w, "Header field %q, Value %q\n", k, v)
    }

}

func main() {
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":5003", nil))
}

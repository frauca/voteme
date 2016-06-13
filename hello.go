package hello

import (
    "fmt"
    "net/http"
)

func init() {
	http.Handle("/static/", http.FileServer(http.Dir(".")))
    http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello, Roger!")
}

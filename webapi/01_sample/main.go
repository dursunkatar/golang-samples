package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Merhaba %s", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Web Server")
	http.ListenAndServe(":8080", nil)
}

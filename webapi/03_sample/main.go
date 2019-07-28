package main

import (
	"html/template"
	"net/http"
)

type Page struct {
	Title           string
	Author          string
	Header          string
	PageDescription string
	Friends         []string
}

func handler(w http.ResponseWriter, r *http.Request) {

	page := Page{
		Title:           "golang",
		Author:          "Dursun Katar",
		Header:          "Bu bir sayfa",
		PageDescription: "Buda sayfanın açıklaması",
		Friends:         []string{"Ramazan", "Mücahit"},
	}
	t, _ := template.ParseFiles("page.html")
	t.Execute(w, page)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func handler(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("index.html")
	tmpl.Execute(w, nil)
}
func handler2(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	value := r.FormValue("deneme")
	value2 := r.Form["deneme"]
	fmt.Println("value1: ", value)
	fmt.Println("value2: ", value2)
	fmt.Fprint(w, "boş")

}
func main() {
	fmt.Println("Web Aktif")
	fmt.Println("Bağlantı için: ip-numara:1453")
	http.HandleFunc("/", handler)
	http.HandleFunc("/parametre", handler2)
	http.ListenAndServe(":1453", nil)
}

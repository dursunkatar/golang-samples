package main

import (
	"net/http"
	"text/template"

	model "./models"
)

var pageTpl, _ = template.ParseFiles("template/page.html")

func handeFunc(w http.ResponseWriter, r *http.Request) {
	coder := getCoder()
	pageTpl.Execute(w, coder)
}

func main() {
	http.HandleFunc("/", handeFunc)
	http.ListenAndServe(":8080", nil)
}

func getCoder() model.Coder {
	coder := model.Coder{
		FirstName: "Dursun",
		LastName:  "Katar",
		Languages: []model.SoftwareLanguage{
			model.SoftwareLanguage{Language: "C#"},
			model.SoftwareLanguage{Language: "Asp.NET MVC"},
			model.SoftwareLanguage{Language: "JavaScript"},
			model.SoftwareLanguage{Language: "Golang"},
		},
	}
	return coder
}

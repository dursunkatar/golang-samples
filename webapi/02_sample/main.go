package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Student struct {
	ID        int    "json:id"
	FirstName string "json:firstname"
	LastName  string "json:lastname"
	Age       int    "json:age"
}
type Info struct {
	Message string "json:message"
}

var apiPath string = "/api"

func main() {
	/*#  /api  #*/
	http.HandleFunc(apiPath, func(w http.ResponseWriter, r *http.Request) {
		message := Info{"/me, /students"}
		response, _ := json.Marshal(message)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, string(response))
	})

	/*#  /api/me  #*/
	http.HandleFunc(apiPath+"/me", func(w http.ResponseWriter, r *http.Request) {
		user := Student{3, "Dursun", "Katar", 28}
		message := user
		response, _ := json.Marshal(message)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, string(response))
	})

	/*#  /api/students  #*/
	http.HandleFunc(apiPath+"/students", func(w http.ResponseWriter, r *http.Request) {
		students := []Student{
			Student{ID: 1, FirstName: "Dursun", LastName: "Katar", Age: 28},
			Student{ID: 2, FirstName: "Ramazan", LastName: "Kaçmaz", Age: 33},
			Student{ID: 3, FirstName: "Mücahit", LastName: "Kaçmaz", Age: 30},
			Student{ID: 3, FirstName: "Ayşe", LastName: "Şen", Age: 30},
		}
		message := students
		response, _ := json.Marshal(message)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, string(response))
	})

	http.ListenAndServe(":8080", nil)
}
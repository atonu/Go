package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func renderTemplate(w http.ResponseWriter, page string) {
	t, err := template.ParseFiles(page)
	if err != nil {
		log.Println(err)
		return
	}
	err = t.Execute(w, nil)
	if err != nil {
		log.Println(err)
		return
	}
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index.html")
}

func test(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "text/html")
	fmt.Fprint(w, "<strong>route made </strong>")
}
func main() {
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/test", test)
	log.Println("Started web server on 8080")
	http.ListenAndServe(":8080", nil)
}

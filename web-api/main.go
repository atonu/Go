package main

import (
	"encoding/json"
	"log"
	"myapp/rps"
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

func playRound(w http.ResponseWriter, r *http.Request) {
	result := rps.PlayRound(1)
	out, err := json.MarshalIndent(result, "", "   ")
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(result)
	w.Header().Set("Content-type", "application/json")
	w.Write(out)
}

func main() {
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/play", playRound)
	log.Println("Started web server on 8080")
	http.ListenAndServe(":8080", nil)
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index.html")
}

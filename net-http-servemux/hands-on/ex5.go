package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.HandlerFunc(ping))
	http.HandleFunc("/pong/", http.HandlerFunc(pong))
	http.HandleFunc("/me/", http.HandlerFunc(identityCatcher))
	http.ListenAndServe(":8080", http.DefaultServeMux)
}

func ping(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Ping!")
}

func pong(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Pong!")
}

func identityCatcher(w http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("home.tpl")
	if err != nil {
		log.Fatalln("error parsing template", err)
	}

	err = tpl.ExecuteTemplate(w, "home.tpl", "Student")
	if err != nil {
		log.Fatalln("error executing template", err)
	}
}

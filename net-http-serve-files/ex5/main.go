package  main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("./templates/index.gohtml"))
}

func index(w http.ResponseWriter, _ *http.Request) {
	if err := tpl.Execute(w, nil); err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("public"))))
	http.Handle("/", http.HandlerFunc(index))

	http.ListenAndServe(":8080", nil)
}
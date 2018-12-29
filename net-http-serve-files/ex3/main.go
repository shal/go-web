package  main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("./starting-files/templates/index.gohtml"))
}

func index(w http.ResponseWriter, _ *http.Request) {
	if err := tpl.Execute(w, nil); err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.Handle("/pics/", http.FileServer(http.Dir("./starting-files/public")))
	http.Handle("/", http.HandlerFunc(index))

	http.ListenAndServe(":8080", nil)
}
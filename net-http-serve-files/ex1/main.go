package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
)

func foo(resp http.ResponseWriter, req *http.Request) {
	io.WriteString(resp, "foo ran")
}

func dogImg(resp http.ResponseWriter, req *http.Request) {
	http.ServeFile(resp, req, "./assets/dog.jpg")
}

func dog(resp http.ResponseWriter, req *http.Request) {
	tpl := template.Must(template.ParseFiles("./dog.gohtml"))
	err := tpl.Execute(resp, req.Header.Get("X-Forwarded-For"))

	if err != nil {
		panic(err)
	}
}

func main() {
	http.Handle("/dog.jpg", http.HandlerFunc(dogImg))
	http.Handle("/dog/", http.HandlerFunc(dog))
	http.Handle("/", http.HandlerFunc(foo))

	http.ListenAndServe(":8080", nil)
}

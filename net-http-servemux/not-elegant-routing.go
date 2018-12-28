package main

import (
	"io"
	"net/http"
)

type handlerImpl string

func (h handlerImpl) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/foo":
		io.WriteString(w, "Foo is here!")
	case "/bar":
		io.WriteString(w, "Bar is here!")
	}
}

func main() {
	var h handlerImpl
	http.ListenAndServe(":8080", h)
}

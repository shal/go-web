package main

import (
	"fmt"
	"net/http"
)

type handlerImpl string

func (h handlerImpl) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "I'm handler!")
}

func main() {
	var handler handlerImpl
	http.ListenAndServe(":8080", handler)
}

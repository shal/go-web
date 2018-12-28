package main

import (
	"fmt"
	"net/http"
)

func c(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(resp, "Hi from c!")
}

func d(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(resp, "Hi from d!")
}

func main() {
	http.HandleFunc("/route1", c)
	http.HandleFunc("/route2", d)

	http.ListenAndServe(":8080", nil)
}
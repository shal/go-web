package main

import (
	"fmt"
	"io"
	"net/http"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	text := fmt.Sprintln("Yeah, I'm on AWS")
	text += fmt.Sprintf("Your public IP address is %s\n", r.Header.Get("X-Forwarded-For"))
	io.WriteString(w, text)
}

func main() {
	http.HandleFunc("/", HelloWorld)
	http.ListenAndServe(":80", nil)
}

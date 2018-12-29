package  main

import "net/http"

func main() {
	handler := http.FileServer(http.Dir("./starting-files"))
	http.ListenAndServe(":8080", handler)
}
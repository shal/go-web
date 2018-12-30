package main

import (
	"log"
	"net/http"
	"strconv"
)

func counter(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("visits")

	 if err == http.ErrNoCookie {
		 http.SetCookie(w, &http.Cookie{
			 Name:  "visits",
			 Value: "0",
		 })
		 return
	 }

	visits, err := strconv.Atoi(cookie.Value)

	if err != nil {
		log.Print(err.Error())
	}

	visits++

	value := strconv.Itoa(visits)

	http.SetCookie(w, &http.Cookie{
		Name:  "visits",
		Value: value,
	})
}

func main() {
	http.HandleFunc("/", counter)

	http.ListenAndServe(":8080", nil)
}

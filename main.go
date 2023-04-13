package main

import (
	"info-center/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/info-center/", infocenterHandler)
	http.ListenAndServe(":8080", nil)
}

func infocenterHandler(w http.ResponseWriter, r *http.Request) {
	topic := r.URL.Path[len("/info-center/"):]

	if r.Method == "POST" {
		handlers.MessageSender(w, r, topic)
	} else if r.Method == "GET" {
		handlers.MessageReceiver(w, r, topic)
	}
}

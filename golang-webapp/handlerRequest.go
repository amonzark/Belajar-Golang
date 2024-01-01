package main

import "net/http"

func handlerIndex(w http.ResponseWriter, r *http.Request) {
	var message = "Welcome"
	w.Write([]byte(message))
}

func handlerHello(w http.ResponseWriter, r *http.Request) {
	var message = "Hello world!"
	w.Write([]byte(message))
}

func handlerPing(w http.ResponseWriter, r *http.Request) {
	var message = "pong"
	w.Write([]byte(message))
}

package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"log"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/callme/ping", PingHandler)
	r.HandleFunc("/callme/error", PingWithRandomError)
	r.HandleFunc("/callme/delay", PingWithRandomDelay)
	log.Fatal(http.ListenAndServe(":8080", r))
}


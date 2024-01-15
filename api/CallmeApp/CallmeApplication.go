package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/callme/ping", PingHandler)
	r.HandleFunc("/callme/error", PingWithRandomError)
	r.HandleFunc("/callme/delay", PingWithRandomDelay)
	log.Fatal(http.ListenAndServe(":8080", r))
}

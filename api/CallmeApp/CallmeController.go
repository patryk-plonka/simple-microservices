package main

import (
	"net/http"
	"math/rand"
	"time"
	"log"
    "strconv"
)

var version string = "1.0"
var name string = "callme"

func PingHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Service: name=%s, version=%s", name, version)
	w.Write([]byte("Service: name=" + name + ", version=" + version))
}

func PingWithRandomError(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UnixNano())
	random := rand.Intn(100)
	if random % 2 == 0 {
		log.Printf("Ping with random error: name=%s, version=%s, random=%d, httpCode=504", name, version, random)
		http.Error(w, "Random error " + strconv.Itoa(random), http.StatusGatewayTimeout)
	} else {
		log.Printf("Ping with random error: name=%s, version=%s, random=%d, httpCode=200", name, version, random)
		w.Write([]byte("Service: name=" + name + ", version=" + version))
	}
}

func PingWithRandomDelay(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UnixNano())
	delay := rand.Intn(3000)
	log.Printf("Ping with random delay: name=%s, version=%s, delay=%d", name, version, delay)
	time.Sleep(time.Duration(delay) * time.Millisecond)
	w.Write([]byte("Service: name=" + name + ", dalay=" + strconv.Itoa(delay)))
}

package main

import (
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	hello := "Hello World!"
	_, err := w.Write([]byte(hello))
	if err != nil {
		log.Print(err)
	}
}

func userAgent(w http.ResponseWriter, r *http.Request) {
	userAgent := r.Header.Get("user-Agent")
	_, err := w.Write([]byte(userAgent))
	if err != nil {
		log.Print(err)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", hello)
	mux.HandleFunc("/agent", userAgent)

	log.Print("running on port 4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}

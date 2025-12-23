package main

import (
	"log"
	"net/http"
)

const addr = ":4000"

func main() {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static"))

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	log.Print("running on port 4000")
	err := http.ListenAndServe(addr, mux)
	log.Fatal(err)
}

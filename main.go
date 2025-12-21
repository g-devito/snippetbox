package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	hello := "Hello World!"
	_, err := w.Write([]byte(hello))
	if err != nil {
		log.Print(err)
	}
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Display a specific snippet..."))
	if err != nil {
		log.Print(err)
	}
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("create a new snippet..."))
	if err != nil {
		log.Print(err)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Print("running on port 4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}

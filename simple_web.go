package main

import (
	"log"
	"net/http"
)

func home(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Welcome"))
}

func main() {
	http.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

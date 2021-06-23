package main

import (
	"github.com/DMax-w/calculator/api"
	"log"
	"net/http"
)

func handleRequests() {
	http.HandleFunc("/api/add", api.AddHandler)
	http.HandleFunc("/api/sub", api.SubHandler)
	http.HandleFunc("/api/mul", api.MulHandler)
	http.HandleFunc("/api/div", api.DivHandler)
	log.Fatal(http.ListenAndServe("", nil))
}

func main() {
	handleRequests()
}

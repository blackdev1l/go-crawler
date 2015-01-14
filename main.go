package main

import (
	"log"
	"net/http"
)

func main() {

	router := NewRouter()
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./www/")))

	//http.Handle("/", http.FileServer(http.Dir("/www")))
	log.Fatal(http.ListenAndServe(":8080", router))

}

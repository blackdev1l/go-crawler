package main

import (
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {

	//router := NewRouter()
	rand.Seed(time.Now().Unix())
	r := mux.NewRouter()
	r.HandleFunc("/", Index)
	r.HandleFunc("/send", Send)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./www/")))
	http.Handle("/", r)

	//http.Handle("/", http.FileServer(http.Dir("/www")))
	log.Fatal(http.ListenAndServe(":8080", r))

}

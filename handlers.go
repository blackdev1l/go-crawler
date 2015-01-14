package main

import (
	"fmt"
	//"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

func loadHtml() []byte {
	fi, err := ioutil.ReadFile("index.html")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(fi))
	return fi

}

func Index(w http.ResponseWriter, r *http.Request) {
	index := loadHtml()
	fmt.Fprint(w, string(index))
}

func Send(w http.ResponseWriter, r *http.Request) {

}

package main

import (
	"fmt"
	//"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

func loadHtml() []byte {
	fi, err := ioutil.ReadFile("www/index.html")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fi)
	return fi

}

func Index(w http.ResponseWriter, r *http.Request) {
	out := loadHtml()
	fmt.Fprintf(w, out)
}

func Send(w http.ResponseWriter, r *http.Request) {

}

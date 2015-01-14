package main

import (
	//"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

type Page struct {
	Title string
	Body  template.HTML
}

/*
func loadHtml() []byte {
	fi, err := ioutil.ReadFile("index.html")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(fi))
	return fi

}
*/

func Index(w http.ResponseWriter, r *http.Request) {
	var page Page
	var templates = template.Must(template.ParseFiles("www/index.html"))
	templates.ExecuteTemplate(w, "index.html", page)
}

func Send(w http.ResponseWriter, r *http.Request) {

}

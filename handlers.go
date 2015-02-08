package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

type Page struct {
	Title string
	Body  template.HTML
}

type UrlResponse struct {
	Url    string `json:"url"`
	Number int    `json:"number"`
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
	log.Println("GET Index")
	var page Page
	var templates = template.Must(template.ParseFiles("www/index.html"))
	templates.ExecuteTemplate(w, "index.html", page)

}

func Send(w http.ResponseWriter, r *http.Request) {

	log.Println("POST Send")
	var url UrlResponse
	g := new(Graph)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(body, &url)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(url.Url))
	log.Println(url.Number)
	g.Url = url.Url

	//Creation of Root Node
	Root := Node{
		Id:    "0",
		X:     len(g.Child) + 1,
		Y:     -1,
		Size:  1,
		Label: url.Url,
	}

	g.Child = append(g.Child, &Root)
	g.Parse(url.Url, url.Number)
	//give to Root Center position
	g.Child[0].X = (len(g.Child) + 1) / 2

	if err := json.NewEncoder(w).Encode(g); err != nil {
		panic(err)
	}
}

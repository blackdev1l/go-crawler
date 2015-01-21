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
	Url string `json:"url"`
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
	n := NewGraph()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	/*
		if err := json.NewEncoder(w).Encode(n); err != nil {
			panic(err)
		}
	*/
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(body, &url)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(url.Url))

	n.Parse(url.Url)
	tmp := new(Node)
	tmp.Url = url.Url
	tmp.X = len(n.Child) + 1
	tmp.Y = 1
	tmp.Label = url.Url
	tmp.Size = 1
	n.Child = append(n.Child, tmp)
	n.Url = ""
	n.X = 0
	n.Y = 0
	if err := json.NewEncoder(w).Encode(n); err != nil {
		panic(err)
	}
}

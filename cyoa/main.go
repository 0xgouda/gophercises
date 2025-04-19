package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"html/template"
	"strings"
)

type Arc struct {
	Title string `json:"title"`
	Story []string `json:"story"`
	Options []struct {
		Text string `json:"text"`
		Name string `json:"arc"`
	}
}

type arcMapper map[string]Arc

func (a *arcMapper) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data, ok := (*a)[strings.TrimLeft(r.URL.Path, "/")]
	if ok {
		tmpl.Execute(w, data)
		return
	}
	tmpl.Execute(w, (*a)["intro"])
}

const startChapter string = "intro" 

var tmpl *template.Template

func main() {
	file, err := os.Open("gopher.json")
	checkErr(err)
	defer file.Close()

	var Arcs arcMapper
	err = json.NewDecoder(file).Decode(&Arcs)
	checkErr(err)

	tmpl = template.Must(template.ParseFiles("layout.html"))

	mux := http.NewServeMux()
	mux.Handle("/", &Arcs)
	
	http.ListenAndServe(":8888", mux)
}

func checkErr(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
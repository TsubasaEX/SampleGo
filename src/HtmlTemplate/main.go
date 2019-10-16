package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type News struct {
	Title string
	Name  string
}

func NewsAggHandle(w http.ResponseWriter, r *http.Request) {
	news := News{"Zach No. 1!!", "Zach"}
	t, _ := template.ParseFiles("basic.html")
	t.Execute(w, news)
}

func IndexHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>index</h1>")
}

func main() {
	http.HandleFunc("/agg", NewsAggHandle)
	http.HandleFunc("/", IndexHandle)
	http.ListenAndServe(":8080", nil)
}

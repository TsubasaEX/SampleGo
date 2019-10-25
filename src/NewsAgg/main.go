package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"text/template"
)

const url = "https://www.washingtonpost.com/news-sitemaps/index.xml"

type NewsLocation struct {
	Location []string `xml:"sitemap>loc"`
}

type News struct {
	Locations []string `xml:"url>loc"`
	Titles    []string `xml:"url>news>title"`
	Keywords  []string `xml:"url>news>keywords"`
}

type NewsMap struct {
	Keyword  string
	Location string
}

type NewsInfo struct {
	Title string
	Map   map[string]NewsMap
}

func NewsAggHandle(w http.ResponseWriter, r *http.Request) {

	resp, _ := http.Get(url)
	bContent, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	var newslocation NewsLocation
	xml.Unmarshal(bContent, &newslocation)

	newsMap := make(map[string]NewsMap)
	for _, loc := range newslocation.Location {
		loc = strings.Replace(loc, "\n", "", -1)
		resp, _ := http.Get(loc)
		bContent, _ := ioutil.ReadAll(resp.Body)
		resp.Body.Close()

		var news News
		xml.Unmarshal(bContent, &news)
		for idx, _ := range news.Keywords {
			newsMap[news.Titles[idx]] = NewsMap{news.Keywords[idx], news.Locations[idx]}
		}
	}

	newsInfo := NewsInfo{"Welcome to Z's News", newsMap}
	t, _ := template.ParseFiles("basic.html")
	fmt.Println(t.Execute(w, newsInfo))
}

func IndexHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>index</h1>")
}

func main() {
	http.HandleFunc("/agg", NewsAggHandle)
	http.HandleFunc("/", IndexHandle)
	http.ListenAndServe(":8080", nil)

}

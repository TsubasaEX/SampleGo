package main

import (
	"encoding/xml"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

const url = "https://www.washingtonpost.com/news-technology-sitemap.xml"

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

	// content := string(bContent)
	// fmt.Println(content)
	var news News
	xml.Unmarshal(bContent, &news)
	// fmt.Println(news)

	// for i := 0; i < len(news.Locations); i++ {
	// 	fmt.Println(news.Locations[i], news.Titles[i], news.Keywords[i])
	// }

	newsMap := make(map[string]NewsMap)
	for idx, _ := range news.Keywords {
		newsMap[news.Titles[idx]] = NewsMap{news.Keywords[idx], news.Locations[idx]}
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

// MoreXmlParser project main.go
package main

import (
	"encoding/xml"
	"fmt"
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

func main() {

	resp, _ := http.Get(url)
	bContent, _ := ioutil.ReadAll(resp.Body)

	// content := string(bContent)
	// fmt.Println(content)
	var news News
	xml.Unmarshal(bContent, &news)
	fmt.Println(news)

	for i := 0; i < len(news.Locations); i++ {
		fmt.Println(news.Locations[i], news.Titles[i], news.Keywords[i])
	}

	newsMap := make(map[string]NewsMap)
	for idx, _ := range news.Titles {
		newsMap[news.Titles[idx]] = NewsMap{news.Keywords[idx], news.Locations[idx]}
	}

	for k, v := range newsMap {
		fmt.Println("Title: ", k)
		fmt.Println("Keyword: ", v.Keyword)
		fmt.Println("Location: ", v.Location)
		fmt.Println("-------------------")
	}
}

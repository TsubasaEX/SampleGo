// XmlParser project main.go
package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
)

// var XmlBytes = []byte(`
// <?xml version="1.0" encoding="UTF-8"?>
// <urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
//    <url>
// 	  <uurl><loc>http://www.example.com/</loc></uurl>
//       <loc>http://www.example.com/</loc>
//       <lastmod>2005-01-01</lastmod>
//       <changefreq>monthly</changefreq>
//       <priority>0.8</priority>
//    </url>
//    <url>
//       <loc>http://www.example.com/catalog?item=12&amp;desc=vacation_hawaii</loc>
//       <changefreq>weekly</changefreq>
//    </url>
//    <url>
//       <loc>http://www.example.com/catalog?item=73&amp;desc=vacation_new_zealand</loc>
//       <lastmod>2004-12-23</lastmod>
//       <changefreq>weekly</changefreq>
//    </url>
//    <url>
//       <loc>http://www.example.com/catalog?item=74&amp;desc=vacation_newfoundland</loc>
//       <lastmod>2004-12-23T18:00:15+00:00</lastmod>
//       <priority>0.3</priority>
//    </url>
//    <url>
//       <loc>http://www.example.com/catalog?item=83&amp;desc=vacation_usa</loc>
//       <lastmod>2004-11-23</lastmod>
//    </url>
// </urlset>
// `)

type Site struct {
	Loc     string `xml:"loc"`
	Lastmod string `xml:"lastmod"`
}

type siteMap struct {
	Sites []Site `xml:"url"`
}

func (s Site) String() string {
	return fmt.Sprintf("%s %s", s.Loc, s.Lastmod)
}

func main() {

	content, err := ioutil.ReadFile("sample.xml")
	if err != nil {
		log.Fatal(err)
	}
	var sm siteMap
	// xml.Unmarshal(XmlBytes, &sm)
	xml.Unmarshal(content, &sm)
	fmt.Println(sm)
	for k, v := range sm.Sites {
		fmt.Println(k, v.Loc, v.Lastmod)
	}

}

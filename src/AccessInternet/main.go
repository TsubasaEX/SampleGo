// AccessInternet project main.go
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

var g string

func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, g)
}

func main() {
	resp, _ := http.Get("https://www.quanzhanketang.com/website/customers.html")
	bytes, _ := ioutil.ReadAll(resp.Body)
	body := string(bytes)
	g = body
	fmt.Println(body)
	fmt.Println(resp.StatusCode)
	resp.Body.Close()
	http.HandleFunc("/", index_handler)
	http.ListenAndServe(":8080", nil)
}

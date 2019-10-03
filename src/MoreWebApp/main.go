// MoreWebApp project main.go
package main

import (
	"fmt"
	"net/http"
)

func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Index</h1>")
}

func about_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1> %s is a idiot </h1>", "Andy")
	fmt.Fprintf(w, "<p>Go is fast!</p>")

	fmt.Fprintf(w, `<h1> I really love Anime!! </h1>
	<h2> I really love exercise!! </h2>
	`)
}

func main() {
	http.HandleFunc("/", index_handler)
	http.HandleFunc("/about", about_handler)
	http.ListenAndServe(":8081", nil)
}

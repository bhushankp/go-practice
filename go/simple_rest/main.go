package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path[1:]
		response := getResponse(path)
		fmt.Fprint(w, response)
	})
	http.ListenAndServe(":8080", nil)

}

func getResponse(path string) string {
	switch path {
	case "hi":
		return "Hello"
	case "how are you":
		return "i'm fine whats about you!!"
	default:
		return "welcome to the go apis"
	}
}

package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "<h1>Hello，这里是 goblog</h1>")
	if err != nil {
		return
	}
}

func main() {
	http.HandleFunc("/", handlerFunc)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}

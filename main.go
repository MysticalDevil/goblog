package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	var err error
	if r.URL.Path == "/" {
		_, err = fmt.Fprintf(w, "<h1>Hello，这里是 goblog</h1>")
	} else if r.URL.Path == "/about" {
		_, err = fmt.Fprintf(w, "此博客为个人练手项目，如有反馈和建议，请联系" + "<a href=\"mailto:devil@gamma@gmail.com\">devil@gamma@gmail.com</a>")
	} else {
		_, err = fmt.Fprintf(w, "<h1>请求页面未找到</h1>")
	}
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

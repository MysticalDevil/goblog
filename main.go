package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	var err error = nil
	if r.URL.Path == "/" {
		_, err = fmt.Fprintf(w, "<h1>Hello，这里是 goblog</h1>")
	} else {
		w.WriteHeader(http.StatusNotFound)
		_, err = fmt.Fprintf(w, "<h1>请求页面未找到 :(</h1>" + "<p>如有疑惑，请联系我</p>")
	}
	if err != nil {
		log.Println(err)
	}
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, err := fmt.Fprintf(w, "此博客为个人练手项目，如有反馈和建议，请联系\n" + "<a href=\"mailto:devil2gamma@gmail.com\">devil2gamma@gmail.com</a>")
	if err != nil {
		log.Println(err)
	}
}

func main() {
	router := http.NewServeMux()

	router.HandleFunc("/", defaultHandler)
	router.HandleFunc("/about", aboutHandler)

	router.HandleFunc("/articles/", func(w http.ResponseWriter, r *http.Request) {
		id := strings.SplitN(r.URL.Path, "/", 3)[2]
		_, _ = fmt.Fprintf(w, "文章 ID：" + id)
	})

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		return
	}
}

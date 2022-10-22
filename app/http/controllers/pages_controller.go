package controllers

import (
	"fmt"
	"net/http"
)

// PagesControllers 处理静态页面
type PagesControllers struct {
}

// Home 首页
func (*PagesControllers) Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello, Welcome to goblog</h1>")
}

// About 管业页面
func (*PagesControllers) About(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "此博客为Demo，如有问题请联系"+"<a href=\"mailto:devil2gamma@gmail.com\">devil2gamma@gmail.com</a>")
}

// NotFound 404 页面
func (*PagesControllers) NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>请求页面未找到</h1><p>如有疑问，请联系我。</p>")
}

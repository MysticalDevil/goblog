package routes

import (
	"github.com/gorilla/mux"
	"goblog/app/http/controllers"
	"net/http"
)

// RegisterWebRoutes 注册网页相关路由
func RegisterWebRoutes(r *mux.Router) {
	// 静态页面
	pc := new(controllers.PagesController)
	r.NotFoundHandler = http.HandlerFunc(pc.NotFound)
	r.HandleFunc("/", pc.Home).Methods("GET").Name("home")
	r.HandleFunc("/about", pc.About).Methods("GET").Name("about")

	// 文章相关页面
	ac := new(controllers.ArticlesController)
	r.HandleFunc("/articles/{id:[0-9]+}", articlesShowHandler).Methods("GET").Name("articles.show")
}

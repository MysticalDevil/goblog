package route

import (
	"github.com/gorilla/mux"
	"goblog/pkg/logger"
	"net/http"
)

// Name2URL 通过路由名称来获取 URL
func Name2URL(routeName string, pairs ...string) string {
	var router *mux.Router
	URL, err := router.Get(routeName).URL(pairs...)
	if err != nil {
		logger.LogError(err)
		return ""
	}
	return URL.String()
}

func GetRouteVariable(parameterName string, r *http.Request) string {
	vars := mux.Vars(r)
	return vars[parameterName]
}

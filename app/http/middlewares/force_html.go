package middlewares

import "net/http"

// ForceHTMLMiddleware 强制标头返回 HTML 内容类型
func ForceHTMLMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 1. 拦截请求并设置标头
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		// 2. 继续处理请求
		next.ServeHTTP(w, r)
	})
}

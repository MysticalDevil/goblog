package view

import (
	"goblog/pkg/logger"
	"goblog/pkg/route"
	"html/template"
	"io"
	"path/filepath"
	"strings"
)

// Render 渲染视图
func Render(w io.Writer, name string, data any) {
	viewDir := "resources/views/"

	// 语法糖 将 articles.show 更正为 articles/show
	name = strings.Replace(name, ".", "/", -1)

	files, err := filepath.Glob(viewDir + "layouts/*.gohtml")
	logger.LogError(err)

	newFiles := append(files,viewDir+name+".gohtml")

	tmpl, err := template.New(name+".gohtml").
		Funcs(template.FuncMap{
			"RouteName2URL": route.Name2URL,
	}).ParseFiles(newFiles...)
	logger.LogError(err)

	err = tmpl.ExecuteTemplate(w, "app", data)
	logger.LogError(err)
}
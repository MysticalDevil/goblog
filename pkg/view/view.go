package view

import (
	"goblog/pkg/auth"
	"goblog/pkg/flash"
	"html/template"
	"io"
	"path/filepath"
	"strings"

	"goblog/pkg/logger"
	"goblog/pkg/route"
)

// D 是 map[string]interface{} 的简写
type D map[string]any

// Render 渲染通用视图
func Render(w io.Writer, data D, tplFiles ...string) {
	RenderTemplate(w, "app", data, tplFiles...)
}

// RenderSimple 渲染简单的视图
func RenderSimple(w io.Writer, data D, tplFiles ...string) {
	RenderTemplate(w, "simple", data, tplFiles...)
}

// RenderTemplate 渲染视图
func RenderTemplate(w io.Writer, name string, data D, tplFiles ...string) {
	data["isLogin"] = auth.Check()
	data["loginUser"] = auth.User
	data["flash"] = flash.All()

	allFiles := getTemplateFiles(tplFiles...)

	tmpl, err := template.New("").
		Funcs(template.FuncMap{
			"RouteName2URL": route.Name2URL,
		}).ParseFiles(allFiles...)
	logger.LogError(err)

	err = tmpl.ExecuteTemplate(w, name, data)
	logger.LogError(err)
}

func getTemplateFiles(tplFiles ...string) []string {
	viewDir := "resources/views/"

	for i, f := range tplFiles {
		tplFiles[i] = viewDir + strings.Replace(f, ".", "/", -1) + ".gohtml"
	}

	layoutFiles, err := filepath.Glob(viewDir + "layouts/*.gohtml")
	logger.LogError(err)

	return append(layoutFiles, tplFiles...)
}

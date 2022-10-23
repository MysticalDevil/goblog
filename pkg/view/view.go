package view

import (
	"goblog/pkg/logger"
	"goblog/pkg/route"
	"html/template"
	"io"
	"path/filepath"
	"strings"
)

type D map[string]any

// Render 渲染视图
func Render(w io.Writer, data any, tplFile ...string) {
	RenderTemplate(w, "app", data, tplFile...)
}

func RenderSimple(w io.Writer, data any, tplFiles ...string) {
	RenderTemplate(w, "simple", data, tplFiles...)
}

func RenderTemplate(w io.Writer, name string, data any, tplFiles ...string) {
	viewDir := "resources/views/"

	// 遍历传参文件列表 Slice，设置正确路径，支持 dir.filename 语法他
	for i, f := range tplFiles {
		tplFiles[i] = viewDir + strings.Replace(f,".", "/", -1) + ".gohtml"
	}

	layoutFiles, err := filepath.Glob(viewDir + "layouts/*.gohtml")
	logger.LogError(err)

	allFiles := append(layoutFiles,tplFiles...)

	tmpl, err := template.New("").
		Funcs(template.FuncMap{
			"RouteName2URL": route.Name2URL,
		}).ParseFiles(allFiles...)
	logger.LogError(err)

	err = tmpl.ExecuteTemplate(w, name, data)
	logger.LogError(err)
}
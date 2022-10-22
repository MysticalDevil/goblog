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
func Render(w io.Writer, data any, tplFile ...string) {
	viewDir := "resources/views/"

	// 遍历传参文件列表 Slice，设置正确路径，支持 dir.filename 语法他
	for i, f := range tplFile {
		tplFile[i] = viewDir + strings.Replace(f,".", "/", -1) + ".gohtml"
	}

	layoutFiles, err := filepath.Glob(viewDir + "layouts/*.gohtml")
	logger.LogError(err)

	allFiles := append(layoutFiles,tplFile...)

	tmpl, err := template.New("").
		Funcs(template.FuncMap{
			"RouteName2URL": route.Name2URL,
	}).ParseFiles(allFiles...)
	logger.LogError(err)

	err = tmpl.ExecuteTemplate(w, "app", data)
	logger.LogError(err)
}
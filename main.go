package main

import (
	"embed"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"goblog/app/http/middlewares"
	"goblog/bootstrap"
	"goblog/config"
	"goblog/pkg/logger"
	"net/http"
)

var router *mux.Router

//go:embed resources/views/articles/*
//go:embed resources/views/auth/*
//go:embed resources/views/categories/*
//go:embed resources/views/layouts/*
var tplFS embed.FS

//go:embed public/*
var staticFS embed.FS

func init() {
	config.Initialize()
}

func main() {
	bootstrap.SetupDB()

	bootstrap.SetupTemplate(tplFS)

	router = bootstrap.SetupRoute(staticFS)

	err := http.ListenAndServe(":8080", middlewares.RemoveTrailingSlash(router))
	logger.LogError(err)
}

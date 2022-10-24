package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"goblog/app/http/middlewares"
	"goblog/bootstrap"
	"goblog/config"
	"goblog/pkg/logger"
	"net/http"
)

var router *mux.Router

func init() {
	config.Initialize()
}

func main() {
	bootstrap.SetupDB()

	router = bootstrap.SetupRoute()

	err := http.ListenAndServe(":8080", middlewares.RemoveTrailingSlash(router))
	logger.LogError(err)
}

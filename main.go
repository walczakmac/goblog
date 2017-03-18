package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"database/sql"
	"html/template"
	"github.com/walczakmac/goblog/model"
)

var connection *sql.DB
var menuItems *[]model.Item

func init() {
	connection = InitialiseDatabaseConnection()
	template.FuncMap{
		"menu": FindMenuItems(connection),
	}
}

func main() {
	defer connection.Close()

	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/css/*filepath", ServeResource)
	router.GET("/js/*filepath", ServeResource)
	router.GET("/images/*filepath", ServeResource)

	http.ListenAndServe(":8080", router)
}

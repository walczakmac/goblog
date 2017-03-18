package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"database/sql"
)

var connection *sql.DB

func init() {
	connection = InitialiseDatabaseConnection()
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

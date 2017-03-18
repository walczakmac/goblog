package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func InitialiseDatabaseConnection() *sql.DB {
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/goblog")
	if err != nil {
		log.Fatalln("Could not connect to database")
	}

	return db
}

func FindAllEntries(connection *sql.DB) *[]Entry {
	res, err := connection.Query("SELECT * FROM entry")
	if err != nil {
		log.Fatalln(err)
	}

	var entries []Entry

	defer res.Close()
	for res.Next() {
		var id int
		var title string
		var content string
		var username string
		var created_at string
		err := res.Scan(&id, &title, &content, &username, &created_at)

		if err != nil {
			log.Fatalln(err)
		}

		entries = append(entries, Entry{id, title, content, username,created_at})
	}

	return &entries
}

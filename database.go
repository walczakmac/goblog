package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"github.com/walczakmac/goblog/model"
	"github.com/walczakmac/goblog/model/menu"
)

func InitialiseDatabaseConnection() *sql.DB {
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/goblog")
	if err != nil {
		log.Fatalln("Could not connect to database")
	}

	return db
}

func FindAllEntries(connection *sql.DB) *[]model.Entry {
	res, err := connection.Query("SELECT * FROM entry")
	if err != nil {
		log.Fatalln(err)
	}

	var entries []model.Entry

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

		entries = append(entries, model.Entry{id, title, content, username,created_at})
	}

	return &entries
}

func FindMenuItems(connection *sql.DB) *[]model.Item {
	res, err := connection.Query("SELECT * FROM menu ORDER BY parent_id ASC")
	if err != nil {
		log.Fatalln(err)
	}

	var items []model.Item

	defer res.Close()

	for res.Next() {
		var id int
		var parent_id sql.NullInt64
		var title string

		err := res.Scan(&id, &parent_id, &title)

		if err != nil {
			log.Fatalln(err)
		}

		item := model.Item{id, parent_id, []model.Item{}, title}
		isChild := false
		for i := 0; i < len(items); i++ {
			if (item.ParentID.Valid && int(item.ParentID.Int64) == items[i].ID) {
				items[i].Children = append(items[i].Children, item)
				isChild = true
				break
			}
		}
		if false == isChild {
			items = append(items, item)
		}
	}

	return &items
}

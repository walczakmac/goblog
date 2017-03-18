package model

import "database/sql"

type Item struct {
	ID int
	ParentID sql.NullInt64
	Children []Item
	Title string
	Url string
}

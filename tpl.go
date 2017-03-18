package main

import (
	"github.com/walczakmac/goblog/model"
	"log"
	"html/template"
)

func GenerateMenu(menuItems *[]model.Item) {
	tpl, err := template.ParseFiles()
	if nil != err {
		log.Fatalln(err)
	}
}
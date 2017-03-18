package main

import (
	"bufio"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"mime"
	"net/http"
	"os"
	"strings"
	"github.com/walczakmac/goblog/model"
	"log"
	"github.com/walczakmac/goblog/model/menu"
)

const baseLayoutPath string = "assets/base_layout/"

var baseLayouts []string = []string{
	baseLayoutPath + "layout.gohtml",
	baseLayoutPath + "header.gohtml",
	baseLayoutPath + "footer.gohtml",
	baseLayoutPath + "sidebar.gohtml",
}

func Index(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	renderTemplate(w, r, "assets/index.gohtml", struct {
		Entries *[]model.Entry
	}{
		FindAllEntries(connection),
	})
}

func ServeResource(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	path := "assets" + req.URL.Path
	pathSplit := strings.Split(path, ".")
	contentType := mime.TypeByExtension("." + pathSplit[len(pathSplit) - 1])

	f, err := os.Open(path)
	if err != nil {
		w.WriteHeader(404)
		f.Close()
	}

	defer f.Close()
	w.Header().Add("Content-type", contentType)
	br := bufio.NewReader(f)
	br.WriteTo(w)
}

func renderTemplate(w http.ResponseWriter, r *http.Request, templateName string, data interface{}) {
	baseLayouts = append(baseLayouts, templateName)
	tpl, err := template.ParseFiles(baseLayouts...)
	if nil != err {
		log.Fatalln(err)
	}

	initialData := struct {
		MenuItems *[]model.Item
		TemplateData interface{}
	}{
		menuItems,
		data,
	}

	tpl.Execute(w, initialData)
}


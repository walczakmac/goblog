package main

import (
	"bufio"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"mime"
	"net/http"
	"os"
	"strings"
)

func Index(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	tpl := template.Must(template.ParseFiles("assets/index.gohtml"))
	tpl.ExecuteTemplate(w, "index.gohtml", struct{
		Entries *[]Entry
	}{
		FindAllEntries(connection),
	})
}

func ServeResource(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	path := "assets" + req.URL.Path
	pathSplit := strings.Split(path, ".")
	contentType := mime.TypeByExtension("." + pathSplit[len(pathSplit)-1])

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

package handler

import (
	"html/template"
	"net/http"
	"path"
)

func HandlerIndex(w http.ResponseWriter, r *http.Request) {
	var filepath = path.Join("views", "index.html")
	var tmpl, err = template.ParseFiles(filepath)
	if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
	}

	var data = map[string]interface{}{
			"title": "Learning Golang Web Map",
			"name":  "Marker Coordinate",
			"modaltitle": "Marker Coordinate",
	}

	err = tmpl.Execute(w, data)
	if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func HandlerGallery(w http.ResponseWriter, r *http.Request) {
	message := "Gallery"
	w.Write([]byte(message))
}
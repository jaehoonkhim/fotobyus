package fotobyus

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

func detailPage(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cfSession")
	data := &imageData{}
	if session.Values["token"] != nil {
		data.Token = session.Values["token"].(string)
	}

	vars := mux.Vars(r)
	data.ImageID = vars["imageID"]
	data.AlbumID = vars["albumID"]

	templates := template.Must(template.ParseFiles("template/detail.tmpl",
		"template/header.tmpl",
		"template/footer.tmpl"))

	err := templates.Execute(w, data)
	if err != nil {
		panic(err)
	}
}

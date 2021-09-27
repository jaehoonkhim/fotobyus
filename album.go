package fotobyus 

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func albumPage(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cfSession")
	data := &albumData{}
	if session.Values["token"] != nil {
		vars := mux.Vars(r)
		data.AlbumID = vars["albumID"]
		data.Token = session.Values["token"].(string)
		templates := template.Must(template.ParseFiles("template/album.tmpl",
			"template/header.tmpl",
			"template/footer.tmpl"))

		err := templates.Execute(w, data)
		if err != nil {
			panic(err)
		}
	} else {
		http.Redirect(w, r, "/", http.StatusFound)
	}

}

func albumSettingsPage(w http.ResponseWriter, r *http.Request) {
	log.Println("albumSettingsPage")
	session, _ := store.Get(r, "cfSession")
	data := &albumData{}
	if session.Values["token"] != nil {
		vars := mux.Vars(r)
		albumID := vars["albumID"]

		data.AlbumID = albumID
		data.Token = session.Values["token"].(string)
		templates := template.Must(template.ParseFiles("template/albumSettings.tmpl",
			"template/header.tmpl",
			"template/footer.tmpl"))

		err := templates.Execute(w, data)
		if err != nil {
			panic(err)
		}
	} else {
		http.Redirect(w, r, "/", http.StatusFound)
	}

}

package fotobyus

import (
	"html/template"
	"net/http"
)

func profilePage(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cfSession")
	data := &tokenData{}
	if session.Values["token"] != nil {
		data.Token = session.Values["token"].(string)

		templates := template.Must(template.ParseFiles("template/profile.tmpl",
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

func profileSettingsPage(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cfSession")
	data := &tokenData{}
	if session.Values["token"] != nil {
		data.Token = session.Values["token"].(string)

		templates := template.Must(template.ParseFiles("template/profileSettings.tmpl",
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

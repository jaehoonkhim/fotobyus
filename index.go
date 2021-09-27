package fotobyus

import (
	"html/template"
	"net/http"
)

func indexPage(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cfSession")
	data := &tokenData{}
	if session.Values["token"] != nil {
		data.Token = session.Values["token"].(string)
	}

	if session.Values["token"] != nil {
		templates := template.Must(template.ParseFiles("template/index.tmpl",
			"template/header.tmpl",
			"template/footer.tmpl"))

		err := templates.Execute(w, data)
		if err != nil {
			panic(err)
		}
	} else {
		http.Redirect(w, r, "/signin", 301)
	}
}

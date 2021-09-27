package fotobyus 

import (
	"html/template"
	"net/http"
)

func guidePage(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cfSession")
	data := &tokenData{}
	if session.Values["token"] != nil {
		data.Token = session.Values["token"].(string)
	}

	templates := template.Must(template.ParseFiles("template/guide.tmpl",
		"template/header.tmpl",
		"template/footer.tmpl"))

	err := templates.Execute(w, data)
	if err != nil {
		panic(err)
	}
}

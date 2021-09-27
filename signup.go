package fotobyus

import (
	"html/template"
	"net/http"
)

func signUpPage(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cfSession")
	token := session.Values["token"]
	if token != nil {
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		templates := template.Must(template.ParseFiles("template/signup.tmpl",
			"template/header.tmpl",
			"template/footer.tmpl"))

		err := templates.Execute(w, nil)
		if err != nil {
			panic(err)
		}
	}
}

func signUpOrgPage(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cfSession")
	token := session.Values["token"]
	if token != nil {
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		templates := template.Must(template.ParseFiles("template/signup_.tmpl",
			"template/header.tmpl",
			"template/footer.tmpl"))

		err := templates.Execute(w, nil)
		if err != nil {
			panic(err)
		}
	}
}

package fotobyus

import (
	"encoding/json"
	"html/template"
	"net/http"
)

func signInPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		session, _ := store.Get(r, "cfSession")
		token := session.Values["token"]
		if token != nil {
			http.Redirect(w, r, "/", http.StatusFound)
		} else {
			templates := template.Must(template.ParseFiles("template/signin.tmpl",
				"template/header.tmpl",
				"template/footer.tmpl"))

			err := templates.Execute(w, nil)
			if err != nil {
				panic(err)
			}
		}
	} else if r.Method == "POST" {
		//json 파라미터 파싱
		type tokenResult struct {
			Status    int    `json:"status"`
			AuthToken string `json:"token"`
		}
		var t tokenResult
		value := r.FormValue("json")
		json.Unmarshal([]byte(value), &t)
		session, _ := store.Get(r, "cfSession")
		session.Values["token"] = t.AuthToken
		session.Save(r, w)

		http.Redirect(w, r, "/profile", http.StatusFound)
	}
}

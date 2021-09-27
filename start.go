package fotobyus

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

func startPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		session, _ := store.Get(r, "cfSession")
		data := &tokenData{}
		if session.Values["token"] != nil {
			data.Token = session.Values["token"].(string)
		}
		if session.Values["token"] == nil {
			http.Redirect(w, r, "/", http.StatusFound)
		} else {
			templates := template.Must(template.ParseFiles("template/start.tmpl",
				"template/header.tmpl",
				"template/footer.tmpl"))

			err := templates.Execute(w, data)
			if err != nil {
				panic(err)
			}
		}
	} else if r.Method == "POST" {
		//json 파라미터 파싱
		type albumResult struct {
			Status  int    `json:"status"`
			AlbumID string `json:"albumid"`
		}
		var album albumResult
		value := r.FormValue("json")
		json.Unmarshal([]byte(value), &album)

		url := fmt.Sprintf("/album/%s", album.AlbumID)
		http.Redirect(w, r, url, http.StatusFound)
	}
}

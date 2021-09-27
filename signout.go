package fotobyus

import "net/http"

type signOutData struct {
	AuthToken string `json:"token"`
}

func signOutPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		session, _ := store.Get(r, "cfSession")
		for key, _ := range session.Values {
			delete(session.Values, key)
		}

		session.Save(r, w)

		http.Redirect(w, r, "/", http.StatusFound)
		// token := session.Values["token"]
		// data := &signOutData{}
		// if token != nil {
		// 	data.AuthToken = token.(string)
		// }
		// pBytes, _ := json.Marshal(data)
		// buff := bytes.NewBuffer(pBytes)
		// http.Post("/api/singnout", "application/json", buff)
	} else if r.Method == "POST" {

	}
}

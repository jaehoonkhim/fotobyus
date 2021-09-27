package api

import (
	"encoding/json"
	"log"
	"net/http"
)

func SignOutHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var data SignoutData

	err := decoder.Decode(&data)
	if err != nil {
		panic(err)
	}

	log.Printf("logout: %s", data.AuthToken)
}

package api

import (
	"encoding/json"
	//"log"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

func SignInHandler(w http.ResponseWriter, r *http.Request) {
	result := &SigninResult{
		Status: 0,
	}

	decoder := json.NewDecoder(r.Body)
	var user User
	err := decoder.Decode(&user)
	if err != nil {
		panic(err)
	}

	c := appengine.NewContext(r)
	key := datastore.NewKey(c, "Users", user.UserID, 0, nil)
	var srchUser User
	err = datastore.Get(c, key, &srchUser)
	if err != nil {
		result.Status = 1
	}

	err = hp.Compare(srchUser.Password, user.Password)
	if srchUser.UserID == user.UserID && err == nil {
		result.Status = 0
		result.AuthToken = key.Encode()
	} else {
		result.Status = 1
	}

	w.Header().Set("Content-Type", "application/json")
	j, jsonErr := json.Marshal(result)
	if jsonErr != nil {
		http.Error(w, jsonErr.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(j)
}

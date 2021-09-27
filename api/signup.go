package api

import (
	"encoding/json"
	//"log"
	"net/http"
	"net/url"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user User
	err := decoder.Decode(&user)
	if err != nil {
		panic(err)
	}

	result := &SignupResult{
		Status: 0,
	}
	email, _ := url.QueryUnescape(user.Email)
	//log.Println("*****api/signuphandler:", email)
	if len(user.UserID) == 0 {
		result.Status = 3
	} else {
		if len(email) == 0 {
			result.Status = 1
		} else if validateEmail(email) == false {
			result.Status = 2
		} else {
			hashpassword, _ := hp.Generate(user.Password)
			user.Password = hashpassword

			c := appengine.NewContext(r)
			key := datastore.NewKey(c, "Users", user.UserID, 0, nil)
			_, putErr := datastore.Put(c, key, &user)
			if putErr != nil {
				http.Error(w, putErr.Error(), http.StatusInternalServerError)
				return
			}
		}
	}

	w.Header().Set("Content-Type", "application/json")
	j, jsonErr := json.Marshal(result)
	if jsonErr != nil {
		http.Error(w, jsonErr.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(j)
}

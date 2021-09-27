package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		postProfile(w, r)
	} else if r.Method == "PUT" {
		putProfile(w, r)
	} else if r.Method == "GET" {
		getProfile(w, r)
	}

}

func postProfile(w http.ResponseWriter, r *http.Request) {

}

func putProfile(w http.ResponseWriter, r *http.Request) {
	log.Println("PutProfile")
	decoder := json.NewDecoder(r.Body)
	var srchUsr ProfileSettingsUser
	err := decoder.Decode(&srchUsr)
	if err != nil {
		panic(err)
	}

	var usr User
	ctx := appengine.NewContext(r)
	key, _ := datastore.DecodeKey(srchUsr.Usertoken)
	err = datastore.Get(ctx, key, &usr)
	if err != nil {
		log.Println(err.Error())
	}
	usr.Username = srchUsr.Username
	usr.Email = srchUsr.Email

	datastore.Put(ctx, key, &usr)

	result := &ResultBase{Status: 0}
	w.Header().Set("Content-Type", "application/json")
	j, jsonErr := json.Marshal(result)
	if jsonErr != nil {
		http.Error(w, jsonErr.Error(), http.StatusInternalServerError)
	}
	w.Write(j)
}

func getProfile(w http.ResponseWriter, r *http.Request) {
	//log.Println("ProfileHandler")
	vars := mux.Vars(r)
	userToken := vars["user"]

	var result ProfileResult

	c := appengine.NewContext(r)
	userKey, _ := datastore.DecodeKey(userToken)
	var user User
	if err := datastore.Get(c, userKey, &user); err != nil {
		panic(err)
	}

	result.Status = 0
	result.Name = user.Username
	result.ID = user.UserID
	result.Email = user.Email

	q := datastore.NewQuery("Albums").Ancestor(userKey)
	var albums []Album
	keys, _ := q.GetAll(c, &albums)

	for i := 0; i < len(albums); i++ {
		var album ProfileAlbum
		album.AlbumID = strconv.FormatInt(keys[i].IntID(), 10)
		album.Title = albums[i].Title
		album.StartDate = albums[i].StartDate
		album.EndDate = albums[i].EndDate
		album.Desc = albums[i].Desc
		result.Albums = append(result.Albums, album)
	}

	w.Header().Set("Content-Type", "application/json")

	j, jsonErr := json.Marshal(result)
	if jsonErr != nil {
		http.Error(w, jsonErr.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(j)

}

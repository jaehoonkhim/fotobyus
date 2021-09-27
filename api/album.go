package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"

	"golang.org/x/net/context"
)

// 앨범을 생성
func AlbumHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		postAlbum(w, r)
	} else if r.Method == "GET" {
		getAlbum(w, r)
	} else if r.Method == "PUT" {
		putAlbum(w, r)
	}
}

func postAlbum(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	decoder := json.NewDecoder(r.Body)
	var album Album
	err := decoder.Decode(&album)
	if err != nil {
		panic(err)
	}

	result := &AlbumResult{
		Status:  0,
		AlbumID: "1",
	}

	//앨범 생성
	parent, _ := datastore.DecodeKey(album.UserToken)
	key := datastore.NewKey(c, "Albums", "", 0, parent)
	albumKey, putErr := datastore.Put(c, key, &album)
	if putErr != nil {
		http.Error(w, putErr.Error(), http.StatusInternalServerError)
		return
	}

	result.AlbumID = strconv.FormatInt(albumKey.IntID(), 10)

	w.Header().Set("Content-Type", "application/json")

	j, jsonErr := json.Marshal(result)
	if jsonErr != nil {
		http.Error(w, jsonErr.Error(), http.StatusInternalServerError)
		return
	}
	log.Println("album json")
	log.Println(string(j))
	w.Write(j)

}

func putAlbum(w http.ResponseWriter, r *http.Request) {
	log.Println(">>>>> putAlbum <<<<<<<<<<<")
	c := appengine.NewContext(r)

	decoder := json.NewDecoder(r.Body)
	var album AlbumUpdate
	var data Album
	err := decoder.Decode(&album)
	if err != nil {
		panic(err)
	}
	data.UserToken = album.UserToken
	data.Title = album.Title
	data.StartDate = album.StartDate
	data.EndDate = album.EndDate
	data.Desc = album.Desc

	log.Printf("%s", album.UserToken)
	log.Printf("%s", album.AlbumID)
	parent, _ := datastore.DecodeKey(album.UserToken)
	key := albumKey(c, album.AlbumID, parent)

	albumKey, _ := datastore.Put(c, key, &data)
	result := &AlbumResult{
		Status:  0,
		AlbumID: "1",
	}
	result.AlbumID = strconv.FormatInt(albumKey.IntID(), 10)
	w.Header().Set("Content-Type", "application/json")

	j, jsonErr := json.Marshal(result)
	if jsonErr != nil {
		http.Error(w, jsonErr.Error(), http.StatusInternalServerError)
		return
	}
	log.Println("album json")
	log.Println(string(j))
	w.Write(j)

}

func albumKey(c context.Context, id string, parent *datastore.Key) *datastore.Key {
	albumID, _ := strconv.ParseInt(id, 10, 64)
	key := datastore.NewKey(c, "Albums", "", albumID, parent)
	return key
}

func getAlbum(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	vars := mux.Vars(r)
	//albumID, _ := strconv.ParseInt(vars["id"], 10, 64)
	userToken := vars["user"]
	userKey, _ := datastore.DecodeKey(userToken)

	//key := datastore.NewKey(c, "Albums", "", albumID, userKey)
	key := albumKey(c, vars["id"], userKey)
	log.Println("stringid: " + key.StringID())
	var album Album
	err := datastore.Get(c, key, &album)
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("album: ", string(album.Title))
	var albumResult = AlbumDataResult{
		Status:    0,
		UserToken: album.UserToken,
		AlbumID:   vars["id"],
		Title:     album.Title,
		StartDate: album.StartDate,
		EndDate:   album.EndDate,
		Desc:      album.Desc,
	}

	w.Header().Set("Content-Type", "application/json")

	j, jsonErr := json.Marshal(albumResult)
	if jsonErr != nil {
		http.Error(w, jsonErr.Error(), http.StatusInternalServerError)
		return
	}
	//log.Println(string(j))
	w.Write(j)

}

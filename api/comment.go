package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

func CommentHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("CommentHandler")
	if r.Method == "POST" {
		postComment(w, r)
	} else if r.Method == "GET" {
		getComment(w, r)
	} else if r.Method == "DELETE" {
		deleteComment(w, r)
	}
}

func postComment(w http.ResponseWriter, r *http.Request) {
	log.Println("postComment")
	decoder := json.NewDecoder(r.Body)
	var comment Comment
	var result CommentResult

	defer writeJson(w, result)

	err := decoder.Decode(&comment)
	if err != nil {
		result.Status = 1
		return
	}

	ctx := appengine.NewContext(r)
	imageKey := imageKey(ctx, comment.ImageID)

	commentKey := datastore.NewKey(ctx, "Comments", "", 0, imageKey)
	datastore.Put(ctx, commentKey, &comment)
}

func getComment(w http.ResponseWriter, r *http.Request) {
	log.Println("getComment")
	vars := mux.Vars(r)
	imageID := vars["id"]
	log.Println(imageID)

	ctx := appengine.NewContext(r)

	imageKey := imageKey(ctx, imageID)
	qry := datastore.NewQuery("Comments").Ancestor(imageKey)
	var comments []Comment
	qry.GetAll(ctx, &comments)

	commentlist := CommentList{Status: 0, Comments: comments}
	w.Header().Set("Content-Type", "application/json")
	j, jsonErr := json.Marshal(commentlist)
	if jsonErr != nil {
		http.Error(w, jsonErr.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(j)
}

func deleteComment(w http.ResponseWriter, r *http.Request) {
	log.Println("deleteComment")
}

func writeJson(w http.ResponseWriter, r CommentResult) {
	w.Header().Set("Content-Type", "application/json")
	j, jsonErr := json.Marshal(r)
	if jsonErr != nil {
		http.Error(w, jsonErr.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(j)
}

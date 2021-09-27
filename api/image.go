package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"net/url"
	"strconv"
	"strings"

	"github.com/gorilla/mux"

	"google.golang.org/appengine"
	"google.golang.org/appengine/blobstore"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/urlfetch"

	"golang.org/x/net/context"
)

func ImageViewHandler(w http.ResponseWriter, r *http.Request) {
	blobstore.Send(w, appengine.BlobKey(r.FormValue("blobKey")))
}

// ImageListHandler 현재 앨범의 전체 이미지 리스트
func ImageListHandler(w http.ResponseWriter, r *http.Request) {
	//log.Println("ImageListHandler")
	vars := mux.Vars(r)
	ctx := appengine.NewContext(r)

	//album id를 이용해서 album key를 구해야한다.
	intID, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		panic(err)
	}
	parent, _ := datastore.DecodeKey(vars["user"])
	albumKey := datastore.NewKey(ctx, "Albums", "", intID, parent)

	qry := datastore.NewQuery("Pics").Ancestor(albumKey)

	var pics []Pic
	qry.GetAll(ctx, &pics)

	pk := &PicKeys{}
	log.Println(len(pics))
	for i := 0; i < len(pics); i++ {
		log.Println(pics[i].PicKey)
		pk.Keys = append(pk.Keys, pics[i].PicKey)
	}

	w.Header().Set("Content-Type", "application/json")
	j, jsonErr := json.Marshal(pk)
	if jsonErr != nil {
		http.Error(w, jsonErr.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(j)
}

var quoteEscaper = strings.NewReplacer("\\", "\\\\", `"`, "\\\"")

func escapeQuotes(s string) string {
	return quoteEscaper.Replace(s)
}

// UploadImageHandler 앨범에 추가되는 사진들을 blobstore에 저장
func UploadImageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		file, fileHeader, _ := r.FormFile("image")
		albumId := r.FormValue("albumid")
		userKey := r.FormValue("userkey")

		//log.Println("albumId:", albumId)
		//log.Println("userkey:", userKey)

		img, _, _ := image.Decode(file)

		var result = UploadImageResult{
			Status: 0,
		}

		c := appengine.NewContext(r)
		albumPath := fmt.Sprintf("album/%s", albumId)
		log.Println(albumPath)
		option := blobstore.UploadURLOptions{
			StorageBucket: albumPath,
		}
		uploadUrl, _ := blobstore.UploadURL(c, "/api/redirect", &option)

		var b bytes.Buffer
		wr := multipart.NewWriter(&b)
		//////createFormImage 구현부...
		h := make(textproto.MIMEHeader)
		h.Set("Content-Disposition",
			fmt.Sprintf(`form-data; name="%s"; filename="%s"`,
				escapeQuotes("image"), escapeQuotes(fileHeader.Filename)))
		h.Set("Content-Type", "image/jpeg")
		fw, err := wr.CreatePart(h)
		////////////////////
		if err != nil {
			result.Status = 1
			w.Header().Set("Content-Type", "application/json")

			j, jsonErr := json.Marshal(result)
			if jsonErr != nil {
				http.Error(w, jsonErr.Error(), http.StatusInternalServerError)
				return
			}
			//log.Println(string(j))
			w.Write(j)
			return
		}

		if err = jpeg.Encode(fw, img, nil); err != nil {
			result.Status = 2
			w.Header().Set("Content-Type", "application/json")

			j, jsonErr := json.Marshal(result)
			if jsonErr != nil {
				http.Error(w, jsonErr.Error(), http.StatusInternalServerError)
				return
			}
			//log.Println(string(j))
			w.Write(j)
			return
		}

		/////////album key////////
		fw, _ = wr.CreateFormField("albumId")
		fw.Write([]byte(albumId))
		///////user key/////////
		fw, _ = wr.CreateFormField("userKey")
		fw.Write([]byte(userKey))
		wr.Close()

		req, err := http.NewRequest("POST", uploadUrl.String(), &b)
		if err != nil {
			result.Status = 3
			w.Header().Set("Content-Type", "application/json")

			j, jsonErr := json.Marshal(result)
			if jsonErr != nil {
				http.Error(w, jsonErr.Error(), http.StatusInternalServerError)
				return
			}
			w.Write(j)
			return
		}

		req.Header.Set("Content-Type", wr.FormDataContentType())
		client := urlfetch.Client(c)
		_, err = client.Do(req)
		if err != nil {
			result.Status = 4
			w.Header().Set("Content-Type", "application/json")

			j, jsonErr := json.Marshal(result)
			if jsonErr != nil {
				http.Error(w, jsonErr.Error(), http.StatusInternalServerError)
				return
			}
			w.Write(j)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		j, jsonErr := json.Marshal(result)
		if jsonErr != nil {
			http.Error(w, jsonErr.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(j)
	}
}

// Redirect blogstore에 저장된 정보를 datastore에 별도로 저장
func Redirect(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	blobs, others, err := blobstore.ParseUpload(r)
	if err != nil {
		log.Println(err)
		return
	}

	file := blobs["image"]
	userKey := others["userKey"]
	albumID := others["albumId"]

	if len(file) == 0 {
		log.Println("no file uploaded")
		return
	}

	log.Printf("len: %d\n", len(file))
	log.Printf("blobkey: %s\n", string(file[0].BlobKey))

	p := Pic{UserKey: userKey[0], AlbumID: albumID[0], PicKey: string(file[0].BlobKey)}

	albumIntID, _ := strconv.ParseInt(albumID[0], 10, 64)
	albumParentKey, _ := datastore.DecodeKey(p.UserKey)
	albumKey := datastore.NewKey(ctx, "Albums", "", albumIntID, albumParentKey)
	key := datastore.NewKey(ctx, "Pics", "", 0, albumKey)
	datastore.Put(ctx, key, &p)
}

func imageKey(ctx context.Context, blobKey string) *datastore.Key {
	picKey, _ := url.QueryUnescape(blobKey)

	qry := datastore.NewQuery("Pics").Filter("PicKey=", picKey).KeysOnly().Limit(1)
	keys, _ := qry.GetAll(ctx, nil)

	return keys[0]
}

func DeleteImageHandler(w http.ResponseWriter, r *http.Request) {
	//log.Println("DeleteImageHandler")
	vars := mux.Vars(r)
	ctx := appengine.NewContext(r)

	blobKey := vars["blobkey"]
	albumID := vars["id"]
	userKey := vars["user"]

	var result = PicResult{Status: 0}

	albumIntID, _ := strconv.ParseInt(albumID, 10, 64)
	albumParentKey, _ := datastore.DecodeKey(userKey)
	albumKey := datastore.NewKey(ctx, "Albums", "", albumIntID, albumParentKey)
	//key := datastore.NewKey(ctx, "Pics", "", 0, albumKey)
	qry := datastore.NewQuery("Pics").Ancestor(albumKey).Filter("PicKey=", blobKey)

	var pic []Pic
	//var picKeys []*datastore.Key
	if picKeys, err := qry.GetAll(ctx, &pic); err != nil {
		result.Status = 1
	} else {
		datastore.Delete(ctx, picKeys[0])
		if err := blobstore.Delete(ctx, appengine.BlobKey(blobKey)); err != nil {
			result.Status = 2
		}

		result.AlbumID = pic[0].AlbumID
	}

	w.Header().Set("Content-Type", "application/json")

	j, jsonErr := json.Marshal(result)
	if jsonErr != nil {
		http.Error(w, jsonErr.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(j)

}

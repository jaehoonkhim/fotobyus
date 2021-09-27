package fotobyus

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"

	"api"
)

var store = sessions.NewCookieStore([]byte("JcEPzkG5y422WEbMnuCg9bCGYGSGJ1VfehuCzZYC7GweiiMUy1E3MdLN2SWBEGYqFpVumbT/cO+FE51dvc9lJw=="))

func init() {

	store.Options = &sessions.Options{
		//Domain:   "www.fotobyus.com",
		Path:     "/",
		MaxAge:   3600 * 1,
		HttpOnly: true,
	}

	r := mux.NewRouter().StrictSlash(false)

	css := http.FileServer(http.Dir("./static/css"))
	img := http.FileServer(http.Dir("./static/img"))
	js := http.FileServer(http.Dir("./static/js"))

	r.Handle("/css/", http.StripPrefix("/css", css))
	r.Handle("/img/", http.StripPrefix("/img", img))
	r.Handle("/js/", http.StripPrefix("/js", js))

	// page
	r.HandleFunc("/", indexPage).Methods("GET")
	r.HandleFunc("/signin", signInPage).Methods("GET")
	r.HandleFunc("/signin", signInPage).Methods("POST")
	r.HandleFunc("/signout", signOutPage).Methods("GET")
	r.HandleFunc("/signout", signOutPage).Methods("POST")
	r.HandleFunc("/signup", signUpPage).Methods("GET")
	r.HandleFunc("/signuporg", signUpOrgPage).Methods("GET")

	r.HandleFunc("/start", startPage).Methods("GET")
	r.HandleFunc("/start", startPage).Methods("POST")
	r.HandleFunc("/album/{albumID}", albumPage).Methods("GET")
	r.HandleFunc("/album/{albumID}/settings", albumSettingsPage).Methods("GET")
	r.HandleFunc("/profile", profilePage).Methods("GET")
	r.HandleFunc("/profile/settings", profileSettingsPage).Methods("GET")
	r.HandleFunc("/detail/{albumID}/{imageID}", detailPage).Methods("GET")

	// api
	// auth
	r.HandleFunc("/api/signin", api.SignInHandler).Methods("POST")
	r.HandleFunc("/api/signup", api.SignUpHandler).Methods("POST")
	r.HandleFunc("/api/album", api.AlbumHandler).Methods("POST")
	r.HandleFunc("/api/signout", api.SignOutHandler).Methods("POST")
	// album
	r.HandleFunc("/api/album", api.AlbumHandler).Methods("GET").Queries("id", "{id}", "user", "{user}")
	r.HandleFunc("/api/album", api.AlbumHandler).Methods("POST")
	r.HandleFunc("/api/album", api.AlbumHandler).Methods("PUT")
	//profile
	r.HandleFunc("/api/profile", api.ProfileHandler).Methods("PUT")
	r.HandleFunc("/api/profile", api.ProfileHandler).Methods("POST")
	r.HandleFunc("/api/profile", api.ProfileHandler).Methods("GET").Queries("user", "{user}")
	// photo
	r.HandleFunc("/api/uploadimage", api.UploadImageHandler).Methods("POST")
	r.HandleFunc("/api/redirect", api.Redirect).Methods("POST")
	r.HandleFunc("/api/image", api.ImageListHandler).Methods("GET").Queries("id", "{id}", "user", "{user}")
	r.HandleFunc("/api/image", api.DeleteImageHandler).Methods("DELETE").Queries("blobkey", "{blobkey}", "id", "{id}", "user", "{user}")
	r.HandleFunc("/api/imageview", api.ImageViewHandler).Methods("GET").Queries("blobKey", "{blobKey}")
	r.HandleFunc("/api/comment", api.CommentHandler).Methods("POST")
	r.HandleFunc("/api/comment", api.CommentHandler).Methods("GET").Queries("id", "{id}")

	http.Handle("/", r)
}

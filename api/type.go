package api

type ResultBase struct {
	Status int `json:"status"`
}

type SignupResult struct {
	Status int `json:"status"`
}

type SigninResult struct {
	Status    int    `json:"status"`
	AuthToken string `json:"token"`
}

type SignoutData struct {
	AuthToken string `json:"token"`
}

type User struct {
	Username string `json:"username"`
	UserID   string `json:"userid"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Album struct {
	UserToken string `json:"user_token"`
	Title     string `json:"title"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	Desc      string `json:"desc"`
}

type AlbumUpdate struct {
	AlbumID   string `json:"album_id"`
	UserToken string `json:"user_token"`
	Title     string `json:"title"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	Desc      string `json:"desc"`
}

type AlbumResult struct {
	Status  int    `json:"status"`
	AlbumID string `json:"albumid"`
}

type AlbumDataResult struct {
	Status    int    `json:"status"`
	UserToken string `json:"user_token"`
	AlbumID   string `json:"albumid"`
	Title     string `json:"title"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	Desc      string `json:"desc"`
}

type ProfileData struct {
	AuthToken string `json:"token"`
}

type ProfileAlbum struct {
	Title     string `json:"title"`
	AlbumID   string `json:"albumid"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	Desc      string `json:"desc"`
}

type ProfileResult struct {
	Status int            `json:"status"`
	Name   string         `json:"name"`
	ID     string         `json:"id"`
	Email  string         `json:"email"`
	Photo  string         `json:"photo"`
	Albums []ProfileAlbum `json:"albums"`
}

type UploadImageResult struct {
	Status    int    `json:"status"`
	ThumbPath string `json:"thumb"`
}

type Pic struct {
	UserKey string `json:"-"`
	AlbumID string `json:"albumid"`
	PicKey  string `json:"-"`
}

type PicKeys struct {
	Keys []string `json:"keys"`
}

type PicResult struct {
	Status  int    `json:"status"`
	AlbumID string `json:"albumid"`
}

type Comment struct {
	ImageID      string `json:"imageid"`
	AlbumID      string `json:"albumid"`
	UserToken    string `json:"usertoken"`
	CommentValue string `json:"commentvalue"`
}

type CommentResult struct {
	Status int `json:"status"`
}

type CommentList struct {
	Status   int       `json:"status"`
	Comments []Comment `json:"comments"`
}

type ProfileSettingsUser struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	Usertoken string `json:"usertoken"`
}

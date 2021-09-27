package api

var hp *HashPassword

func init() {
	hp = NewHashPassword("shrinklabs.com")
}

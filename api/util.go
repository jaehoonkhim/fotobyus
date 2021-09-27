package api

import "regexp"

func validateEmail(email string) bool {
	r := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return r.MatchString(email)
}

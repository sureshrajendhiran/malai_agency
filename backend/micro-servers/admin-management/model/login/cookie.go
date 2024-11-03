package login

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/securecookie"
)

var hashKey = []byte("!#%KSASFRS@1532")
var s = securecookie.New(hashKey, nil)

// SetCookieHandler func used to set cookie
func SetCookieHandler(r *http.Request, w http.ResponseWriter, userId string, value string) {
	cookieValue := ReadCookieHandler(r)
	cookieValue[userId] = value
	encoded, err := s.Encode("x-cookie", cookieValue)
	expiration := time.Now().Add(365 * 24 * time.Hour)
	// expiration := time.Now().Add(0 * 10 * time.Second)
	// Domain:  "localhost",
	if err == nil {
		cookie := &http.Cookie{
			Name:    "x-cookie",
			Value:   encoded,
			Expires: expiration,
			Path:    "/",
		}
		http.SetCookie(w, cookie)
	}
}

func ReadCookieHandler(r *http.Request) map[string]interface{} {
	value := make(map[string]interface{})
	cookie, err := r.Cookie("x-cookie")
	if err != nil {
		log.Println("error in read cookie", err)
		return value
	}
	err = s.Decode("x-cookie", cookie.Value, &value)
	if err != nil {
		log.Println("error in decode cookie", err)
		return value
	}
	return value
}

// CheckCookieHistory func used to check login user detail and cookie is same or not
func CheckCookieHistory(r *http.Request, loginUserID string) bool {
	// cookie name used to get cookie
	cookie, err := r.Cookie("x-cookie")
	if err != nil {
		log.Println("CheckCookieHistory error in read cookie", err)
		return true
	}
	// Decode cookie string to map format
	value := make(map[string]interface{})
	err = s.Decode("x-cookie", cookie.Value, &value)
	if err != nil {
		log.Println("CheckCookieHistory error in decode cookie", err)
		return true
	}
	// Check current id present exiting cookie or not
	if value[loginUserID] == nil || value[loginUserID] == "0" {
		return true
	}
	return false
}

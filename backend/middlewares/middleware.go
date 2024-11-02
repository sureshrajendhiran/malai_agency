package middlewares

import (
	"back_end_servers/services/auth"
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

// LoggingMiddleware func used to check token and validate
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Do stuff here
		if !strings.Contains(r.RequestURI, "login") &&
			!strings.Contains(r.RequestURI, "/admin/appInfo/") {
			if !auth.Before(w, r) {
				return
			}
		}
		log.Println(r.RequestURI, r.Header["user_id"])
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

func JsonValid(req *http.Request) error {
	decoder := json.NewDecoder(req.Body)
	data := make(map[string]interface{})
	err := decoder.Decode(&data)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

package apiserver

import (
	"crypto/subtle"
	"encoding/json"
	"net/http"
	"os"
)

var (
	username, password string
)

func init() {
	username = os.Getenv("AUTH_USERNAME")
	if username == "" {
		username = "test"
	}

	password = os.Getenv("AUTH_PASSWORD")
	if password == "" {
		password = "test"
	}
}

func BasicAuth(w http.ResponseWriter, r *http.Request, realm string) bool {
	user, pass, ok := r.BasicAuth()

	hasValidUsername := subtle.ConstantTimeCompare([]byte(user), []byte(username)) != 1
	hasValidPassword := subtle.ConstantTimeCompare([]byte(pass), []byte(password)) != 1

	if !ok || hasValidPassword || hasValidUsername {
		w.Header().Set("WWW-Authenticate", `Basic realm="`+realm+`"`)
		json.NewEncoder(w).Encode(NewError(w, r, 401, "Unauthorized"))

		return false
	}

	return true
}

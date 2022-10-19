package web

import (
	"net/http"
)

type basicAuthHandler struct {
	//next http.Handler
	next func(w http.ResponseWriter, r *http.Request)
}

func (b *basicAuthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	username, password, ok := r.BasicAuth()
	if ok {

		if isUserAthorized(username, password) {
			setUseNameHeader(r, username)
			b.next(w, r)
			return
		}
	}

	w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
	http.Error(w, "Unauthorized", http.StatusUnauthorized)

}

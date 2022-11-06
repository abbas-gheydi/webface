package web

import (
	"net/http"
)

type coockieAuthHandler struct {
	//next http.Handler
	next func(w http.ResponseWriter, r *http.Request)
}

func (h *coockieAuthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	token, readCookieError := r.Cookie("auth")

	if readCookieError == nil && isCookieValied(token.Value) {
		//authenticatin is successful
		username, _ := validate_jwt(token.Value)
		setUseNameHeader(r, username)

		h.next(w, r)

	} else {

		http.Redirect(w, r, "/login/", http.StatusFound)

	}

}

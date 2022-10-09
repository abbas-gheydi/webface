package web

import "net/http"

type authHandler struct {
	//next http.Handler
	next func(w http.ResponseWriter, r *http.Request)
}

func (h *authHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	token, readCookieError := r.Cookie("auth")

	if readCookieError == nil && isCookieValied(token.Value) {
		//authenticatin is successful
		h.next(w, r)

	} else {

		http.Redirect(w, r, "/login/", http.StatusFound)

	}

}

func MustAuth(handler func(w http.ResponseWriter, r *http.Request)) *authHandler {

	return &authHandler{next: handler}
}

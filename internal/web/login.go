package web

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
)

var jwtHmacSecret []byte

func init() {
	generateRandomString := func() string {
		randomByte := make([]byte, 8)
		rand.Read(randomByte)
		hash := md5.Sum([]byte(randomByte))
		return hex.EncodeToString(hash[:])

	}

	jwtHmacSecret = []byte(generateRandomString())

}

func setCookie(w http.ResponseWriter, uname string) {
	cookie := &http.Cookie{Name: "auth",
		Value: generate_jwt(uname),
		Path:  "/",
	}
	http.SetCookie(w, cookie)

}

func isCookieValied(postedCookie string) bool {
	if postedCookie == "" {
		return false
	} else {
		_, status := validate_jwt(postedCookie)
		return status

	}
}
func generate_jwt(user string) string {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": user,
		"exp":  time.Now().Add(time.Hour * 8).Unix(),
	})

	tokenString, _ := token.SignedString(jwtHmacSecret)

	//fmt.Println(tokenString, err)
	return tokenString
}
func validate_jwt(tokenString string) (user string, isValied bool) {

	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return jwtHmacSecret, nil
	})
	if token == nil {
		return "", false

	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		//fmt.Println(claims["user"], claims["nbf"])
		user = fmt.Sprint(claims["user"])
		return user, true
	} else {
		//log.Println(err)
		return "", false
	}

}

type app struct {
	Name, Desc template.HTML
}

func getAppNameAndDesc() app {

	return app{
		Name: "Demo App",
		Desc: "please login with domain credential",
	}
}

func loginPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		t, err := template.ParseFiles("web/login.html")
		//t, err := template.New("login").Parse(loginTemplate)

		if err != nil {
			log.Println(err)
			return

		}

		execute_err := t.Execute(w, getAppNameAndDesc())
		if execute_err != nil {

			log.Println(execute_err)
		}
	}
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")
		if isUserAthorized(username, password) {

			setCookie(w, username)
			http.Redirect(w, r, "/", http.StatusFound)
			return

		}
		http.Redirect(w, r, "/login/", http.StatusFound)
	}
}

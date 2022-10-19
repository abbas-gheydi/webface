package web

import (
	"net/http"

	"github.com/Abbas-gheydi/webface/internal/auth"
	ldapAuth "github.com/korylprince/go-ad-auth/v3"
)

const (
	basic_auth = "basic_auth"
)

var (
	LdapServer          string
	LdapPort            = 389
	LdapSecurityLevel   = 4
	LdapBaseDN          string
	LdapGroup           string
	AUTH_MODE           string
	LISTEN_ADDR         string = "0.0.0.0:8080"
	USERNAME_HEADER     string = "X-username"
	SET_USERNAME_HEADER bool
)

var SSO authSource

func MustAuth(handler func(w http.ResponseWriter, r *http.Request)) http.Handler {
	if AUTH_MODE == basic_auth {
		return &basicAuthHandler{next: handler}

	}

	return &coockieAuthHandler{next: handler}
}

type authSource interface {
	IsUserAuthenticated(username string, password string, group string) (authStat bool)
}

func SetAuthSource(source string) authSource {
	if source == "ldap" {

		return auth.LdapProvider{
			LdapConfig: &ldapAuth.Config{
				Server:   LdapServer,
				Port:     LdapPort,
				Security: ldapAuth.SecurityType(LdapSecurityLevel),
				BaseDN:   LdapBaseDN,
			},
		}
	}
	return nil
}

func isUserAthorized(usernmae string, password string) bool {

	return SSO.IsUserAuthenticated(usernmae, password, LdapGroup)
}

func setUseNameHeader(r *http.Request, username string) {
	if SET_USERNAME_HEADER {
		r.Header.Add(USERNAME_HEADER, username)
	}

}

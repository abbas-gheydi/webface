package web

import (
	"net/http"

	"github.com/Abbas-gheydi/webface/internal/auth"
	ldapAuth "github.com/korylprince/go-ad-auth/v3"
)

const (
	basic_auth = "basic_auth"
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
		r.Header.Set(USERNAME_HEADER, username)
	}

}

func setBearerTokenHeader(r *http.Request) {
	if UPSTREAM_Bearer_TOKEN != "" {
		r.Header.Set("Authorization", "Bearer "+UPSTREAM_Bearer_TOKEN)
	}
}

func setK8sDashboardBearerToken(r *http.Request, username string) {
	if EnableK8sDashbaord {

		r.Header.Set("Authorization", "Bearer "+auth.GetK8sToken(K8sSaNameSpace, username))
	}
}

func setRequestExtraHeaders(r *http.Request, username string) {
	setUseNameHeader(r, username)
	setBearerTokenHeader(r)
	setK8sDashboardBearerToken(r, username)

}

package web

import (
	"github.com/Abbas-gheydi/webface/internal/auth"
	ldapAuth "github.com/korylprince/go-ad-auth/v3"
)

var (
	LdapServer        string
	LdapPort          = 389
	LdapSecurityLevel = 4
	LdapBaseDN        string
	LdapGroup         string
)

var SSO authSource

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

package main

import (
	"os"
	"strconv"

	"github.com/Abbas-gheydi/webface/internal/web"
)

func init() {

	web.UpStream = os.Getenv("UPSTREAM")
	web.LdapServer = os.Getenv("LDAP_SERVER")
	web.LdapPort, _ = strconv.Atoi(os.Getenv("LDAP_PORT"))
	web.LdapSecurityLevel, _ = strconv.Atoi(os.Getenv("LDAP_SEC_LEVEL"))
	web.LdapBaseDN = os.Getenv("LDAP_BASEDN")
	web.LdapGroup = os.Getenv("LDAP_GROUP")
	web.AUTH_MODE = os.Getenv("AUTH_MODE")
	web.SET_USERNAME_HEADER, _ = strconv.ParseBool(os.Getenv("SET_USERNAME_HEADER"))

	if os.Getenv("LISTEN_ADDR") != "" {
		web.LISTEN_ADDR = os.Getenv("LISTEN_ADDR")
	}

	if web.LdapPort == 0 {
		web.LdapPort = 389
	}

	web.SSO = web.SetAuthSource("ldap")

}
func main() {

	web.Router()
}

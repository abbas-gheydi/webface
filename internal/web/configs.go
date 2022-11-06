package web

var (
	LdapServer            string
	LdapPort              = 389
	LdapSecurityLevel     = 4
	LdapBaseDN            string
	LdapGroup             string
	AUTH_MODE             string
	LISTEN_ADDR           string = "0.0.0.0:8080"
	USERNAME_HEADER       string = "X-username"
	SET_USERNAME_HEADER   bool
	InsecureSkipVerify    bool
	UPSTREAM_Bearer_TOKEN string
)

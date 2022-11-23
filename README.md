## About The Project:

"webFace" is an authentication proxy server that protects services that do not have authentication mechanisms.   
It authenticates users against Microsoft's active directory using the ldap/ldaps protocol.
It has two authenticate mechanism:   
basic auth   
web form    
to configure it you can use these envs:  
```

UPSTREAM='http://url'
LDAP_SERVER='ldap server ip'
LDAP_PORT='389'
LDAP_SEC_LEVEL='4'
LDAP_BASEDN='DC=test,DC=local'
LDAP_GROUP='groupname'
SET_USERNAME_HEADER='false'
AUTH_MODE="basic_auth"
InsecureSkipTLSVerify=false
UPSTREAM_Bearer_TOKEN=""
K8S_DASHBOARD_MODE=false

```


# Curl examples for Janssen

Curl-based examples help show the function of the API without getting bogged down in whether the golang is wrong


## secrets



## UMA API

```
source env.secret
curl -s -X GET \
"https://${host}/.well-known/uma2-configuration 
```

Get a token
```
curl -k -u '0008-b52a8524-35b2-4835-968e-481a366be8cd:TVtZwLZxp25XFDelMJNDQsa8' \
    https://demoexample.gluu.org/oxauth/restv1/token \
    -d grant_type=client_credentials
```

```
GET http://sample.com/oxauth/restv1/host/rsrc/resource_set HTTP/1.1
Authorization: Bearer ${token}
```

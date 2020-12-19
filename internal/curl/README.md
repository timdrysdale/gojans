# Curl examples for Janssen

Trying out some commands with curl helps gain familiarity with the API. You can do this as follows.

## Set up Jans

this one is over to you - see auth server (setup instructions](https://github.com/JanssenProject/jans-auth-server).

## Try it out

First, export some common variables

```
export host=jans.your.host
export client_id=p9l9adf...
export client_secret=425jlk345....
```
Or if you want to do this more than one session, you can automate it as follows by editing the example secret exports script

``` 
cd ${gojans}/internal/curl
cp export.example export.secret
emacs export.secret 
source export.secret
```

Files with the extension ```.secret``` are ignored in this repo, so your credentials will not be added.

### Client (RP) token

For your test client, make sure it has ```client_credentials``` grant type enabled iun the openID Connect Client configuration. It takes a minute or two to take effect from when it the configuration is updated. For details of the request, see the [token endpoint details](https://gluu.org/docs/gluu-server/4.2/api-guide/openid-connect-api/).

To see the full details of the request which is sent, you can add ```--trace-ascii /dev/stdout``` e.g.

```
curl --trace-ascii /dev/stdout -X  POST \
	 -H "Content-Type: application/x-www-form-urlencoded" \
    -d "grant_type=client_credentials&client_id=${client_id}&client_secret=${client_secret}"\
    "https://${host}/oxauth/restv1/token" \
```

To store the token in a variable, use ```jq``` to parse the response

```
client_token=$(curl -s -X POST\
	 -H "Content-Type: application/x-www-form-urlencoded" \
    -d "grant_type=client_credentials&client_id=${client_id}&client_secret=${client_secret}"\
    "https://${host}/oxauth/restv1/token" \
	| jq -r '.access_token')
echo $client_token	
```
It takes about a minute changes to the client_credentials grant status to be reflected in getting a token or not.

## UMA API

TODO - fix this section

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

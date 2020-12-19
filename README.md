# gojans
(Possibly) a Go adapter for [Janssen (Gluu) auth server](https://github.com/JanssenProject/home)

![alt text][status]

## Rationale
Things changed, so here we are. Ancient history: After reviewing the options in late November 2020, I started contributing to [gocloak](https://github.com/Nerzal/gocloak) in anticipation of adopting [keycloak](https://www.keycloak.org/) for my [new remote laboratory infrastructure](https://practable.io). Just a few weeks later in early Dec 2020, there was shift in Red Hat's FOSS policy, i.e. [dropping CENTOS](https://arstechnica.com/gadgets/2020/12/centos-shifts-from-red-hat-unbranded-to-red-hat-beta/). It seems likely that [Keycloak is at significant risk of the same fate](https://www.gluu.org/blog/keycloak-is-the-next-centos/). If you were still considering [keycloak](https://www.keycloak.org/) then there is a [good golang adapter available](https://github.com/Nerzal/gocloak) which offers features not found in [goth](https://github.com/markbates/goth) because they are intended for different tasks (building own systems vs logging into others'). 

Now I would like something like ```gocloak``` but for the [Janssen Project](https://github.com/JanssenProject/home), which in short form is known as ```jans```. There is an OpenAPI spec for ```janssen``` which means ```codegen``` (aka swagger in a previous life) might be all that is needed. This repo is my exploration of ```jans``` in relation to golang, and will only turn into a repo if ```codegen``` doesn't do what we need here.

## openAPI

For handling openAPI3.0, there is [kin-openapi](https://github.com/getkin/kin-openapi). It can generate schemas with [openapi3gen]( https://godoc.org/github.com/getkin/kin-openapi/openapi3gen), but not client code. That can be done with [oapi-codegen](https://github.com/deepmap/oapi-codegen).

## openAPI 3.0

Importing the oxAuth openAPI3.0 description into apimatic throws one warning on import:
```
Endpoint group Token contains an endpoint get-introspection whose parameter response_as_jwt example value is not valid. Parameter removed from the generated testcase.
```
Then 9 warnings on validation:
```
Model ClientResponse has a field name run_introspection_script_before_access_token_as_jwt_creation_and_include_claims that may be too long. Beware, this may cause issues in Code Generation.
Model IntrospectionResponse contains invalid example value for field iat. Invalid format: Input string was not in a correct format. Ignoring this example.
Model RegisterParams contains invalid example value for field jwks. Invalid format: Error reading JArray from JsonReader. Current JsonReader item is not an array: String. Path '', line 1, position 67. Ignoring this example.
Model RegisterParams has a field name run_introspection_script_before_access_token_as_jwt_creation_and_include_claims that may be too long. Beware, this may cause issues in Code Generation.
Model UmaResource contains invalid example value for field resource_scopes. Missing bracket(s) for an array. Ignoring this example.
Endpoint get-introspection contains invalid example value for parameter response_as_jwt. Invalid format: String was not recognized as a valid Boolean. Ignoring this example.
Endpoint put-host-rsrc-resource_set{rsid} has an unused template parameter named rsid.
Endpoint delete-host-rsrc-resource_set has an unused template parameter named rsid.
Endpoint get-uma_scopes has an unused template parameter named id.
One or more elements in the API specification has a missing description field.
```

AFter some edits, the warnings are reduced to
```
5 warnings
Model ClientResponse has a field name run_introspection_script_before_access_token_as_jwt_creation_and_include_claims that may be too long. Beware, this may cause issues in Code Generation.
Model RegisterParams has a field name run_introspection_script_before_access_token_as_jwt_creation_and_include_claims that may be too long. Beware, this may cause issues in Code Generation.
Endpoint put-host-rsrc-resource_set{rsid} has an unused template parameter named rsid.
Endpoint delete-host-rsrc-resource_set has an unused template parameter named rsid.
Endpoint get-uma_scopes has an unused template parameter named id.
1 message
One or more elements in the API specification has a missing description field.
```

Exporting as openAPI2.0 in yaml format, go-swagger throws this error
```
unexpected map key type, got: bool
```
This is a [known issue](https://github.com/go-swagger/go-swagger/issues/1209) with 'n' being a reserved word in YAML for boolean false.
```
{ Y, true, Yes, ON  }    : Boolean true
{ n, FALSE, No, off }    : Boolean false
```

This line can be found by looking for isolated keys named n or y, such as with
```
cat oxAuthModifiedApimaticOpenAPI20.yaml | grep -n '^[[:blank:]]*n:'
```

This also threw up a 'y' which needed corrected a few lines below.

```
$ swagger validate oxAuthModifiedApimaticOpenAPI20.yaml
2020/12/19 20:53:31 
The swagger spec at "oxAuthModifiedApimaticOpenAPI20.yaml" is valid against swagger specification 2.0
2020/12/19 20:53:31 
The swagger spec at "oxAuthModifiedApimaticOpenAPI20.yaml" showed up some valid but possibly unwanted constructs.
2020/12/19 20:53:31 See warnings below:
2020/12/19 20:53:31 - WARNING: example value for body in body does not validate its schema
2020/12/19 20:53:31 - WARNING: body.software_version.example in body must be of type string: "number"
2020/12/19 20:53:31 - WARNING: in operation "session_status", example value in response 200 does not validate its schema
2020/12/19 20:53:31 - WARNING: 200.auth_time.example in body must be of type date: "float64"
2020/12/19 20:53:31 - WARNING: in operation "get-register", example value in response 200 does not validate its schema
2020/12/19 20:53:31 - WARNING: 200.software_version.example in body must be of type string: "number"
2020/12/19 20:53:31 - WARNING: definitions.ClientResponse.software_version.example in body must be of type string: "number"
2020/12/19 20:53:31 - WARNING: definitions.RegisterParams.software_version.example in body must be of type string: "number"
2020/12/19 20:53:31 - WARNING: definitions.SessionStateObject.auth_time.example in body must be of type date: "float64"
```

The date was fixed to integer/int32 similar to other authorisation time instances.

The swagger now passes validation.

### generation

```
swagger generate client -f oxAuthModifiedApimaticOpenAPI20.yaml -A oxAuth 
```

```
swagger generate client -f oxAuthModifiedApimaticOpenAPI20.yaml -A oxAuth \
--additional-initialism=UMA   \
--additional-initialism=FIDO  \
--additional-initialism=UMA2  \
--additional-initialism=FIDO2 \
--additional-initialism=U2F   \
--additional-initialism=JWK   \
--additional-initialism=JWKS  
```

Then install dependencies
```
go get -u -f ./...
```


### openAPI 3.0  to openAPI 2.0 conversion

The openAPI swagger file is [here](https://raw.githubusercontent.com/JanssenProject/jans-auth-server/master/docs/oxAuthSwagger.yaml) and can be converted to openAPI2.0 [here](https://apimatic.io) with a free account.

The conversion throws a number of messages

`
8 warnings:
Endpoint group Token contains an endpoint get-introspection whose parameter response_as_jwt example value is not valid. Parameter removed from the generated testcase.
Model IntrospectionResponse contains invalid example value for field iat. Invalid format: Input string was not in a correct format. Ignoring this example.
Model RegisterParams contains invalid example value for field jwks. Invalid format: Error reading JArray from JsonReader. Current JsonReader item is not an array: String. Path '', line 1, position 67. Ignoring this example.
Model UmaResource contains invalid example value for field resource_scopes. Missing bracket(s) for an array. Ignoring this example.
Endpoint get-introspection contains invalid example value for parameter response_as_jwt. Invalid format: String was not recognized as a valid Boolean. Ignoring this example.
Endpoint put-host-rsrc-resource_set{rsid} has an unused template parameter named rsid.
Endpoint delete-host-rsrc-resource_set has an unused template parameter named rsid.
Endpoint get-uma_scopes has an unused template parameter named id.
21 messages
Input identified as OpenAPI v3.0 (YAML).
A test case generated from example values for endpoint get_clientinfo in group Client Info.
A test case generated from example values for endpoint well-known-gluu-configuration in group Server Configuration.
A test case generated from example values for endpoint get-introspection in group Token.
A test case generated from example values for endpoint jwks in group JWK - JSON Web Key Set (JWKs).
A test case generated from example values for endpoint post-register in group Registration.
A test case generated from example values for endpoint session_status in group Session Management.
A test case generated from example values for endpoint end_session in group Session Management.
A test case generated from example values for endpoint get-uma-gather_claims in group UMA (User Managed Access).
A test case generated from example values for endpoint post-uma-gather_claims in group UMA (User Managed Access).
A test case generated from example values for endpoint uma2-configuration in group UMA (User Managed Access).
A test case generated from example values for endpoint post-fido-u2f-authentication in group FIDO U2F.
A test case generated from example values for endpoint fido-configuration in group FIDO U2F.
A test case generated from example values for endpoint post-fido-u2f-registration in group FIDO U2F.
A test case generated from example values for endpoint get-fido2-configuration in group FIDO2.
A test case generated from example values for endpoint options in group FIDO2.
A test case generated from example values for endpoint result in group FIDO2.
A test case generated from example values for endpoint attestation-options in group FIDO2.
A test case generated from example values for endpoint attestation-result in group FIDO2.
Imported: 43 Endpoints, 44 Models.
One or more elements in the API specification has a missing description field.
Proceed
``

```


Some inconsistencies or issues are detected by the [go-swagger](https://github.com/go-swagger/go-swagger) validation, so the 2.0 yaml needed hand modifying, which has been done with the aim of passing validation and some inconsistencies may remain. Examples of changes:

  - arrays of string examples needed fixing
  - date time format
  - JWKS format (both in schema upload and getting keys when jwks_uri not an option)
  
## Curl

Trying out some commands with curl helps check the modified spec is ok (or not!) and gain familiarity with the API. You can do this as follows.

export some common variables

```
export host=jans.your.host
export client_id=p9l9adf...
export client_secret=425jlk345....
```

Make sure the client has the ```client_credentials``` grant type enabled iun the openID Connect Client configuration. It takes a minute or two to take effect from when it the configuration is updated. For details of the request, see the [token endpoint details](https://gluu.org/docs/gluu-server/4.2/api-guide/openid-connect-api/).

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

Query... how quickly will permissions added to the server be accessible?!
  
gluu ser


## Mocking

This can be done automatically with [apisprout](https://github.com/danielgtaylor/apisprout) although it would require the examples to be set usefully.

[status]: https://img.shields.io/badge/status-concept-yellow "concept status"

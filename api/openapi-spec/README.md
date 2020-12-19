# jansAuth swagger

The openAPI3.0 spec for the ```jans-auth``` server which is the same as for the oxAuth server in the Gluu project.

This required some conversion (to swagger 2.0) and some editing (to fix some inconsistencies).

## openAPI3.0 and golang

Swagger's openAPI descendent ```swagger-codegen``` can handle openAPI3.0 but produced a flat implementation when tried. This is clunky - for example, all optional fields included as mandatory parameters in function calls. Alternatives for handling openAPI3.0, there is [kin-openapi](https://github.com/getkin/kin-openapi). It can generate schemas with [openapi3gen]( https://godoc.org/github.com/getkin/kin-openapi/openapi3gen), but not client code. That can be done with [oapi-codegen](https://github.com/deepmap/oapi-codegen), although it does not favour strongly typed code because it prioritises simplicity. 

A mature interface is available in [go-swagger](https://github.com/go-swagger/go-swagger/), which has concentrated on producing good code against ```swagger 2.0``` specs. 

## openAPI 3.0 editing and conversion

The openAPI3.0 spec for ```jans-auth-server``` is [here](https://raw.githubusercontent.com/JanssenProject/jans-auth-server/master/docs/oxAuthSwagger.yaml).

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

After exporting as openAPI2.0 in yaml format, validating in with go-swagger throws this error

```
$ swagger validate jansAuth.yaml
  unexpected map key type, got: bool
```
This is a [known issue](https://github.com/go-swagger/go-swagger/issues/1209) with 'n' and 'y' being a reserved word in YAML for boolean false and true, but also being keys in JWK.
```
{ Y, true, Yes, ON  }    : Boolean true
{ n, FALSE, No, off }    : Boolean false
```

This line can be found by looking for isolated keys named n or y, such as with
```
cat oxAuthModifiedApimaticOpenAPI20.yaml | grep -n '^[[:blank:]]*n:'
```
This also threw up a 'y' which needed corrected a few lines below. The fix is to put the keys in quotes to hint they are strings not bools.

The date was fixed to integer/int32 similar to other authorisation time instances.

The swagger now passes validation.


### generation

The code is generated as follows, including some initialisms which we want the generator to keep intact (to avoid names like ```u_m_a_2...```)

```
swagger generate client -f jansAuth.yaml -A jansAuth \
--additional-initialism=UMA   \
--additional-initialism=FIDO  \
--additional-initialism=UMA2  \
--additional-initialism=FIDO2 \
--additional-initialism=U2F   \
--additional-initialism=JWK   \
--additional-initialism=JWKS
```

Once generation is complete, the dependencies can be installed
```
go get -u -f ./...
```

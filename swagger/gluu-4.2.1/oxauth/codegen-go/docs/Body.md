# Body

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Scope** | **string** | OpenID Connect requests MUST contain the openid scope value. If the openid scope value is not present, the behavior is entirely unspecified. Other scope values MAY be present. | [default to null]
**ResponseType** | **string** | OAuth 2.0 Response Type value that determines the authorization processing flow to be used, including what parameters are returned from the endpoints used. | [default to null]
**ClientId** | **string** | OAuth 2.0 Client Identifier valid at the Authorization Server. | [default to null]
**RedirectUri** | **string** | Redirection URI to which the response will be sent. This URI MUST exactly match one of the Redirection URI values for the Client pre-registered at the OpenID Provider. | [default to null]
**State** | **string** | Opaque value used to maintain state between the request and the callback. | [optional] [default to null]
**ResponseMode** | **string** | Informs the Authorization Server of the mechanism to be used for returning parameters from the Authorization Endpoint. | [optional] [default to null]
**Nonce** | **string** | String value used to associate a Client session with an ID Token, and to mitigate replay attacks. | [optional] [default to null]
**Display** | **string** | ASCII string value that specifies how the Authorization Server displays the authentication and consent user interface pages to the End-User. | [optional] [default to null]
**Prompt** | **string** | Space delimited, case sensitive list of ASCII string values that specifies whether the Authorization Server prompts the End-User for reauthentication and consent. | [optional] [default to null]
**MaxAge** | **int32** | Maximum Authentication Age. Specifies the allowable elapsed time in seconds since the last time the End-User was actively authenticated by the OP. | [optional] [default to null]
**UiLocales** | **string** | End-User&#x27;s preferred languages and scripts for the user interface, represented as a space-separated list of BCP47 [RFC5646] language tag values, ordered by preference. | [optional] [default to null]
**IdTokenHint** | **string** | ID Token previously issued by the Authorization Server being passed as a hint about the End-User&#x27;s current or past authenticated session with the Client. If the End-User identified by the ID Token is logged in or is logged in by the request, then the Authorization Server returns a positive response. | [optional] [default to null]
**LoginHint** | **string** | Hint to the Authorization Server about the login identifier the End-User might use to log in (if necessary). | [optional] [default to null]
**AcrValues** | **string** | Requested Authentication Context Class Reference values. Space-separated string that specifies the acr values that the Authorization Server is being requested to use for processing this Authentication Request, with the values appearing in order of preference. | [optional] [default to null]
**AmrValues** | **string** | AMR Values. | [optional] [default to null]
**Request** | **string** | This parameter enables OpenID Connect requests to be passed in a single, self-contained parameter and to be optionally signed and/or encrypted. The parameter value is a Request Object value. It represents the request as a JWT whose Claims are the request parameters. | [optional] [default to null]
**RequestUri** | **string** | This parameter enables OpenID Connect requests to be passed by reference, rather than by value. The request_uri value is a URL using the https scheme referencing a resource containing a Request Object value, which is a JWT containing the request parameters. | [optional] [default to null]
**RequestSessionId** | **string** | Request session id. | [optional] [default to null]
**SessionId** | **string** | Session id of this call. | [optional] [default to null]
**OriginHeaders** | **string** | Origin headers. Used in custom workflows. | [optional] [default to null]
**CodeChallenge** | **string** | PKCE code challenge. | [optional] [default to null]
**CodeChallengeMethod** | **string** | PKCE code challenge method. | [optional] [default to null]
**CustomResponseHeaders** | **string** | Custom Response Headers. | [optional] [default to null]
**Claims** | **string** | Requested Claims. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


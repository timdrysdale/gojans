# Body7

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GrantType** | **[]string** | Provide a list of the OAuth 2.0 grant types that the Client is declaring that it will restrict itself to using. | [default to null]
**Code** | **string** | Code which is returned by authorization endpoint. (For grant_type&#x3D;authorization_code) | [optional] [default to null]
**RedirectUri** | **string** | Redirection URI to which the response will be sent. This URI MUST exactly match one of the Redirection URI values for the Client pre-registered at the OpenID Provider. | [optional] [default to null]
**Username** | **string** | End-User username. | [optional] [default to null]
**Password** | **string** | End-User password. | [optional] [default to null]
**Scope** | **[]string** | OpenID Connect requests MUST contain the openid scope value. If the openid scope value is not present, the behavior is entirely unspecified. Other scope values MAY be present. Scope values used that are not understood by an implementation SHOULD be ignored. | [optional] [default to null]
**Assertion** | **string** | Assertion. | [optional] [default to null]
**RefreshToken** | **string** | Refresh token. | [optional] [default to null]
**ClientId** | **string** | OAuth 2.0 Client Identifier valid at the Authorization Server. | [optional] [default to null]
**ClientSecret** | **string** | The client secret.  The client MAY omit the parameter if the client secret is an empty string. | [optional] [default to null]
**CodeVerifier** | **string** | The client&#x27;s PKCE code verifier. | [optional] [default to null]
**Ticket** | **string** |  | [optional] [default to null]
**ClaimToken** | **string** |  | [optional] [default to null]
**ClaimTokenFormat** | **string** |  | [optional] [default to null]
**Pct** | **string** |  | [optional] [default to null]
**Rpt** | **string** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


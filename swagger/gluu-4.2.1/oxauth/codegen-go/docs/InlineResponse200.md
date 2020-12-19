# InlineResponse200

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | **string** | The access token issued by the authorization server. | [default to null]
**TokenType** | **string** | The access token type provides the client with the information required to successfully utilize the access token to make a protected resource request (along with type-specific attributes). | [default to null]
**ExpiresIn** | **int32** | The lifetime in seconds of the access token. For example, the value \\\&quot;3600\\\&quot; denotes that the access token will expire in one hour from the time the response was generated. | [optional] [default to null]
**RefreshToken** | **string** | The refresh token, which can be used to obtain new access tokens using the same authorization grant | [optional] [default to null]
**Scope** | **[]string** |  | [optional] [default to null]
**IdToken** | **string** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


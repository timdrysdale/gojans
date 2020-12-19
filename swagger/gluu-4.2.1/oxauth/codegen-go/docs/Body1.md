# Body1

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | **string** | OAuth 2.0 Client Identifier valid at the Authorization Server. | [default to null]
**Scope** | **string** | CIBA authentication requests must contain the openid scope value. | [default to null]
**ClientNotificationToken** | **string** | It is a bearer token provided by the Client that will be used by the OpenID Provider to authenticate the callback request to the Client. It is required if the Client is registered to use Ping or Push modes. | [default to null]
**AcrValues** | **string** | Requested Authentication Context Class Reference values. | [optional] [default to null]
**LoginHintToken** | **string** | A token containing information identifying the end-user for whom authentication is being requested. | [optional] [default to null]
**IdTokenHint** | **string** | An ID Token previously issued to the Client by the OpenID Provider being passed back as a hint to identify the end-user for whom authentication is being requested. | [optional] [default to null]
**LoginHint** | **string** | A hint to the OpenID Provider regarding the end-user for whom authentication is being requested. | [optional] [default to null]
**BindingMessage** | **string** | A human readable identifier or message intended to be displayed on both the consumption device and the authentication device to interlock them together for the transaction by way of a visual cue for the end-user. | [optional] [default to null]
**UserCode** | **string** | A secret code, such as password or pin, known only to the user but verifiable by the OP. | [optional] [default to null]
**RequestedExpiry** | **int32** | A positive integer allowing the client to request the expires_in value for the auth_req_id the server will return. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


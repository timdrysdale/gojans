# RegisterResponseParam

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | **string** | Unique Client Identifier. It MUST NOT be currently valid for any other registered Client. | [default to null]
**ClientSecret** | **string** | This value is used by Confidential Clients to authenticate to the Token Endpoint | [optional] [default to null]
**RegistrationAccessToken** | **string** | Registration Access Token that can be used at the Client Configuration Endpoint to perform subsequent operations upon the Client registration. | [optional] [default to null]
**RegistrationClientUri** | **string** | Location of the Client Configuration Endpoint where the Registration Access Token can be used to perform subsequent operations upon the resulting Client registration. | [optional] [default to null]
**ClientIdIssuedAt** | **int32** | Time at which the Client Identifier was issued. | [optional] [default to null]
**ClientSecretExpiresAt** | **int32** | Time at which the client_secret will expire or 0 if it will not expire. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


# AuthenticateRequestMessage

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | **string** | Version of the protocol that the to-be-registered U2F token must speak. | [default to null]
**Challenge** | **string** | The websafe-base64-encoded challenge. | [default to null]
**AppId** | **string** | The application id that the RP would like to assert. | [default to null]
**KeyHandle** | **string** | Websafe-base64 encoding of the key handle obtained from the U2F token during registration. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


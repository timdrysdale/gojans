# RptIntrospectionResponse1

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | **bool** | Boolean indicator of whether or not the presented token is currently active. | [default to null]
**Exp** | **int64** | Integer timestamp, in seconds since January 1 1970 UTC, indicating when this token will expire. | [optional] [default to null]
**Iat** | **int32** | Integer timestamp, measured in the number of seconds since January 1 1970 UTC, indicating when this permission was originally issued. | [optional] [default to null]
**ClientId** | **string** | Client id used to obtain RPT. | [optional] [default to null]
**Sub** | **string** | Subject of the token. Usually a machine-readable identifier of the resource owner who authorized this token. | [optional] [default to null]
**Aud** | **string** | Service-specific string identifier or list of string identifiers representing the intended audience for this token. | [optional] [default to null]
**Permissions** | [**[]RptIntrospectionResponsePermissions**](RptIntrospectionResponse_permissions.md) |  | [default to null]
**PctClaims** | **map[string]string** |  | [optional] [default to null]
**Iss** | **string** | String representing the issuer of this token, as defined in JWT [RFC7519]. | [optional] [default to null]
**Jti** | **string** | String identifier for the token, as defined in JWT [RFC7519]. | [optional] [default to null]
**Nbf** | **int32** | Integer timestamp, measured in the number of seconds since January 1 1970 UTC, indicating the time before which this permission is not valid. | [optional] [default to null]
**ResourceId** | **string** | Resource ID. | [default to null]
**ResourceScopes** | **[]string** |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


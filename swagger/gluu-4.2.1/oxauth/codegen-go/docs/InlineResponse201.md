# InlineResponse201

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceId** | **string** | The identifier for a resource to which this client is seeking access. The identifier MUST correspond to a resource that was previously registered. | [default to null]
**ResourceScopes** | **[]string** | An array referencing zero or more strings representing scopes to which access was granted for this resource. Each string MUST correspond to a scope that was registered by this resource server for the referenced resource. | [default to null]
**Params** | **map[string]string** | A key/value map that can contain custom parameters. | [optional] [default to null]
**Exp** | **int64** | Number of seconds since January 1 1970 UTC, indicating when this token will expire. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


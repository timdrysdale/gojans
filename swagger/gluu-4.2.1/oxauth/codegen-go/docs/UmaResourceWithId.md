# UmaResourceWithId

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | UMA Resource identifier | [default to null]
**Name** | **string** | A human-readable string describing a set of one or more resources. This name MAY be used by the authorization server in its resource owner user interface for the resource owner. | [optional] [default to null]
**Uri** | **string** | A human-readable string describing the resource | [optional] [default to null]
**Type_** | **string** | A string uniquely identifying the semantics of the resource set. For example, if the resource set consists of a single resource that is an identity claim that leverages standardized claim semantics for \\\&quot;verified email address\\\&quot;, the value of this property could be an identifying URI for this claim. | [optional] [default to null]
**Scopes** | **[]string** | An array of strings, any of which MAY be a URI, indicating the available scopes for this resource set. URIs MUST resolve to scope descriptions as defined in Section 2.1. Published scope descriptions MAY reside anywhere on the web; a resource server is not required to self-host scope descriptions and may wish to point to standardized scope descriptions residing elsewhere. It is the resource server&#x27;s responsibility to ensure that scope description documents are accessible to authorization servers through GET calls to support any user interface requirements. The resource server and authorization server are presumed to have separately negotiated any required interpretation of scope handling not conveyed through scope descriptions. | [optional] [default to null]
**ScopeExpression** | **string** |  | [optional] [default to null]
**Description** | **string** | A human-readable string describing the resource | [optional] [default to null]
**IconUri** | **string** | A URI for a graphic icon representing the resource set. The referenced icon MAY be used by the authorization server in its resource owner user interface for the resource owner. | [optional] [default to null]
**Iat** | **int64** | number of seconds since January 1 1970 UTC, indicating when the token was issued at | [default to null]
**Exp** | **int64** | number of seconds since January 1 1970 UTC, indicating when this token will expire. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


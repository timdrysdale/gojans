# {{classname}}

All URIs are relative to *https://virtserver.swaggerhub.com/yuriyz1/oxauth/4.2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetIntrospection**](TokenApi.md#GetIntrospection) | **Get** /introspection | The Introspection OAuth 2 Endpoint.
[**PostIntrospection**](TokenApi.md#PostIntrospection) | **Post** /introspection | The Introspection OAuth 2 Endpoint.
[**PostToken**](TokenApi.md#PostToken) | **Post** /token | To obtain an Access Token, an ID Token, and optionally a Refresh Token, the RP (Client).
[**Revoke**](TokenApi.md#Revoke) | **Post** /revoke | Revoke an Access Token or a Refresh Token, the RP (Client).

# **GetIntrospection**
> IntrospectionResponse GetIntrospection(ctx, authorization, token, optional)
The Introspection OAuth 2 Endpoint.

The Introspection OAuth 2 Endpoint.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorization** | **string**| Client Authorization details that contains the access token along with other details. | 
  **token** | **string**|  | 
 **optional** | ***TokenApiGetIntrospectionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TokenApiGetIntrospectionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **tokenTypeHint** | **optional.String**| ID Token previously issued by the Authorization Server being passed as a hint about the End-User. | 
 **responseAsJwt** | **optional.Bool**| OPTIONAL. Boolean value with default value false. If true, returns introspection response as JWT (signed based on client configuration used for authentication to Introspection Endpoint). | 

### Return type

[**IntrospectionResponse**](IntrospectionResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostIntrospection**
> IntrospectionResponse PostIntrospection(ctx, authorization, optional)
The Introspection OAuth 2 Endpoint.

The Introspection OAuth 2 Endpoint.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorization** | **string**| Client Authorization details that contains the access token along with other details. | 
 **optional** | ***TokenApiPostIntrospectionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TokenApiPostIntrospectionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **token** | **optional.**|  | 

### Return type

[**IntrospectionResponse**](IntrospectionResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostToken**
> InlineResponse200 PostToken(ctx, grantType, code, redirectUri, username, password, scope, assertion, refreshToken, clientId, clientSecret, codeVerifier, ticket, claimToken, claimTokenFormat, pct, rpt)
To obtain an Access Token, an ID Token, and optionally a Refresh Token, the RP (Client).

To obtain an Access Token, an ID Token, and optionally a Refresh Token, the RP (Client).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **grantType** | [**[]string**](string.md)|  | 
  **code** | **string**|  | 
  **redirectUri** | **string**|  | 
  **username** | **string**|  | 
  **password** | **string**|  | 
  **scope** | [**[]string**](string.md)|  | 
  **assertion** | **string**|  | 
  **refreshToken** | **string**|  | 
  **clientId** | **string**|  | 
  **clientSecret** | **string**|  | 
  **codeVerifier** | **string**|  | 
  **ticket** | **string**|  | 
  **claimToken** | **string**|  | 
  **claimTokenFormat** | **string**|  | 
  **pct** | **string**|  | 
  **rpt** | **string**|  | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Revoke**
> Revoke(ctx, token, tokenTypeHint)
Revoke an Access Token or a Refresh Token, the RP (Client).

Revoke an Access Token or a Refresh Token, the RP (Client).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**|  | 
  **tokenTypeHint** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: content, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


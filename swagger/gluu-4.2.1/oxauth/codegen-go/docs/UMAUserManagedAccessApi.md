# {{classname}}

All URIs are relative to *https://virtserver.swaggerhub.com/yuriyz1/oxauth/4.2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUmaGatherClaims**](UMAUserManagedAccessApi.md#GetUmaGatherClaims) | **Get** /uma/gather_claims | UMA Claims Gathering Endpoint.
[**HostRsrcPr**](UMAUserManagedAccessApi.md#HostRsrcPr) | **Post** /host/rsrc_pr | Registers permission.
[**PostUmaGatherClaims**](UMAUserManagedAccessApi.md#PostUmaGatherClaims) | **Post** /uma/gather_claims | UMA Claims Gathering Endpoint
[**Uma2Configuration**](UMAUserManagedAccessApi.md#Uma2Configuration) | **Get** /uma2-configuration | Gets UMA configuration data.

# **GetUmaGatherClaims**
> GetUmaGatherClaims(ctx, optional)
UMA Claims Gathering Endpoint.

UMA Claims Gathering Endpoint.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UMAUserManagedAccessApiGetUmaGatherClaimsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UMAUserManagedAccessApiGetUmaGatherClaimsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **optional.String**| OAuth 2.0 Client Identifier valid at the Authorization Server. | 
 **ticket** | **optional.String**|  | 
 **claimsRedirectUri** | **optional.String**|  | 
 **state** | **optional.String**|  | 
 **reset** | **optional.Bool**|  | 
 **authentication** | **optional.Bool**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HostRsrcPr**
> []InlineResponse201 HostRsrcPr(ctx, resourceId, resourceScopes, params, authorization)
Registers permission.

Registers permission.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **resourceId** | **string**|  | 
  **resourceScopes** | [**[]string**](string.md)|  | 
  **params** | [**map[string]string**](string.md)|  | 
  **authorization** | **string**| Client Authorization details that contains the access token along with other details. | 

### Return type

[**[]InlineResponse201**](inline_response_201.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostUmaGatherClaims**
> PostUmaGatherClaims(ctx, optional)
UMA Claims Gathering Endpoint

UMA Claims Gathering Endpoint

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UMAUserManagedAccessApiPostUmaGatherClaimsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UMAUserManagedAccessApiPostUmaGatherClaimsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **optional.**|  | 
 **ticket** | **optional.**|  | 
 **claimsRedirectUri** | **optional.**|  | 
 **state** | **optional.**|  | 
 **reset** | **optional.**|  | 
 **authentication** | **optional.**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Uma2Configuration**
> InlineResponse2001 Uma2Configuration(ctx, )
Gets UMA configuration data.

Gets UMA configuration data.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


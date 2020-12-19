# {{classname}}

All URIs are relative to *https://virtserver.swaggerhub.com/yuriyz1/oxauth/4.2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteHostRsrcResourceSet**](UMA2ResourceApi.md#DeleteHostRsrcResourceSet) | **Delete** /host/rsrc/resource_set | Deletes a previously registered resource.
[**GetHostRsrcResourceSet**](UMA2ResourceApi.md#GetHostRsrcResourceSet) | **Get** /host/rsrc/resource_set | Lists all previously registered resource.
[**GetHostRsrcResourceSetrsid**](UMA2ResourceApi.md#GetHostRsrcResourceSetrsid) | **Get** /host/rsrc/resource_set/{rsid} | Reads a previously registered resource.
[**PostHostRsrcResourceSet**](UMA2ResourceApi.md#PostHostRsrcResourceSet) | **Post** /host/rsrc/resource_set | Adds a new resource description.
[**PutHostRsrcResourceSetrsid**](UMA2ResourceApi.md#PutHostRsrcResourceSetrsid) | **Put** /host/rsrc/resource_set | Updates a previously registered resource.

# **DeleteHostRsrcResourceSet**
> DeleteHostRsrcResourceSet(ctx, authorization, rsid)
Deletes a previously registered resource.

Deletes a previously registered resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorization** | **string**|  | 
  **rsid** | **string**| Resource ID | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHostRsrcResourceSet**
> []string GetHostRsrcResourceSet(ctx, authorization, optional)
Lists all previously registered resource.

Lists all previously registered resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorization** | **string**|  | 
 **optional** | ***UMA2ResourceApiGetHostRsrcResourceSetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UMA2ResourceApiGetHostRsrcResourceSetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **scope** | **optional.String**| Scope uri. | 

### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHostRsrcResourceSetrsid**
> UmaResourceWithId GetHostRsrcResourceSetrsid(ctx, authorization, rsid)
Reads a previously registered resource.

Reads a previously registered resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorization** | **string**| Client Authorization details that contains the access token along with other details. | 
  **rsid** | **string**| Resource description ID. | 

### Return type

[**UmaResourceWithId**](UmaResourceWithId.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostHostRsrcResourceSet**
> UmaResourceResponse PostHostRsrcResourceSet(ctx, authorization, optional)
Adds a new resource description.

Adds a new resource description.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorization** | **string**| Client Authorization details that contains the access token along with other details. | 
 **optional** | ***UMA2ResourceApiPostHostRsrcResourceSetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UMA2ResourceApiPostHostRsrcResourceSetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of UmaResource1**](UmaResource1.md)|  | 

### Return type

[**UmaResourceResponse**](UmaResourceResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutHostRsrcResourceSetrsid**
> UmaResourceResponse PutHostRsrcResourceSetrsid(ctx, authorization, rsid, optional)
Updates a previously registered resource.

Updates a previously registered resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorization** | **string**|  | 
  **rsid** | **string**| Resource ID. | 
 **optional** | ***UMA2ResourceApiPutHostRsrcResourceSetrsidOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UMA2ResourceApiPutHostRsrcResourceSetrsidOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of UmaResource**](UmaResource.md)|  | 

### Return type

[**UmaResourceResponse**](UmaResourceResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

